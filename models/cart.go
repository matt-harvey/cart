package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/matt-harvey/cart/db"
)

type Cart struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Items     CartItems `json:"items" has_many:"cart_items"`
}

// QualifiesFor returns (true, nil) only if there are Items *persisted* for this
// Cart, such that they cause it to qualify for the passed Promotion.
func (c *Cart) QualifiesFor(promotion Promotion) (bool, error) {
	var result []uint
	err := db.Conn().RawQuery(
		"SELECT SUM(quantity) FROM cart_items WHERE cart_id = ? AND product_id = ?",
		c.ID,
		promotion.RequiredProductID).All(&result)
	if err != nil {
		return false, err
	}
	sum := result[0]
	return (sum >= promotion.RequiredProductQuantity), nil
}

// ApplyPricing calculates which Promotions apply to this cart, and sets the StandardPriceCents and
// DiscountedPriceCents of its contained Items accordingly. This assumes Items have already been
// persisted, and will load them into memory from the database if they haven't been loaded already.
// NOTE This method does *not* persist the Items after their pricing is mutated.
func (c *Cart) ApplyPricing() error {
	err := db.Conn().Eager("Items.Product").Find(c, c.ID)
	if err != nil {
		return err
	}
	var itemProductIDs []int
	for i, item := range c.Items {
		standardPriceCents := uint(item.Product.PriceCents) * item.Quantity
		c.Items[i].StandardPriceCents = standardPriceCents   // mutate via index
		c.Items[i].DiscountedPriceCents = standardPriceCents // we'll revisit this below
		itemProductIDs = append(itemProductIDs, item.ProductID)
	}
	// Find the Promotions where that *might* apply to this Cart (because their RequiredProductID
	// corresponds to a Product that's in this Cart).
	shortlistedPromotions := []Promotion{}
	// Apparent	inability of binding slice into statement using Pop, means we're stuck building a
	// query fragment string manually that looks like "1, 5, 7" (the itemProductIDs). It's safe though
	// because we know they're ints! Horrible, but at least safe.
	// TODO Put this in a utility function somewhere.
	queryFragment := strings.Trim(strings.Join(strings.Split(fmt.Sprint(itemProductIDs), " "), ", "), "[]")
	queryFragment = fmt.Sprintf("required_product_id IN (%s)", queryFragment)
	err = db.Conn().Where(queryFragment).All(&shortlistedPromotions)
	if err != nil {
		return err
	}

	// Apply each qualifying Promotion to the Cart.
	for _, promotion := range shortlistedPromotions {
		qualifies, err := c.QualifiesFor(promotion)
		if err != nil {
			return err
		}
		if !qualifies {
			continue
		}
		discount, err := promotion.Discount()
		if err != nil {
			return err
		}
		var discountedProductIDs []int
		// TODO This probably belongs in Promotion model.
		// TODO Should also filter to only find promotion_discounted_products where their
		// product_id is that of one of the Products in this Cart (to avoid loading more
		// rows than necessary). We could reuse the utility function for interpolating IDs slice
		// alluded to above.
		err = db.Conn().RawQuery(
			"SELECT product_id FROM promotion_discounted_products WHERE promotion_id = ?",
			promotion.ID).
			All(&discountedProductIDs)
		if err != nil {
			return err
		}
		for _, discountedProductID := range discountedProductIDs {
			for i, item := range c.Items {
				if item.ProductID == discountedProductID {
					// FIXME Avoid applying it if it's one of the required items (e.g. if
					// every pair of socks after the second one is half-price, then we need
					// to respect the "after" part of that).
					c.Items[i].ApplyDiscount(discount)
				}
			}
		}
	}

	// FIXME
	return nil
}
