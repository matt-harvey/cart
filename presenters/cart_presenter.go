package presenters

import (
	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

type CartPresenter struct {
	ID              int                `json:"id"`
	TotalPriceCents uint               `json:"total_price_cents"`
	Items           CartItemPresenters `json:"items"`
}

func NewCartPresenter(cartID int) (*CartPresenter, error) {
	cartPresenter := CartPresenter{ID: cartID}
	cart := models.Cart{}
	err := db.Conn().Eager("Items.Product").Find(&cart, cartID)
	if err != nil {
		return nil, err
	}
	err = cart.ApplyPricing()
	if err != nil {
		return nil, err
	}
	// Remember which Products we already have Items for in this Cart, so we can merge them
	// into a single line in the presenter. For each Product ID, we remember the position (index)
	// in the CartPresenter.Items at which this Product appears.
	productIDsCovered := make(map[int]int)

	// TODO DRY this up.
	for _, cartItem := range cart.Items {
		pos, ok := productIDsCovered[cartItem.Product.ID]
		if ok {
			// We've already got a CartItemPresenter in CartPresenter.Items for this Product;
			// so let's merge this CartItem's Quantity and prices into the existing CartItemPresenter.
			cartPresenter.Items[pos].Quantity += cartItem.Quantity
			cartPresenter.Items[pos].StandardPriceCents += cartItem.StandardPriceCents
			cartPresenter.Items[pos].DiscountCents +=
				(cartItem.StandardPriceCents - cartItem.DiscountedPriceCents)
			cartPresenter.Items[pos].DiscountedPriceCents += cartItem.DiscountedPriceCents
		} else {
			cartItemPresenter := CartItemPresenter{
				ProductID:            cartItem.Product.ID,
				ProductName:          cartItem.Product.Name,
				Quantity:             cartItem.Quantity,
				StandardPriceCents:   cartItem.StandardPriceCents,
				DiscountCents:        cartItem.StandardPriceCents - cartItem.DiscountedPriceCents,
				DiscountedPriceCents: cartItem.DiscountedPriceCents,
			}
			productIDsCovered[cartItem.Product.ID] = len(cartPresenter.Items)
			cartPresenter.Items = append(cartPresenter.Items, cartItemPresenter)
		}
		cartPresenter.TotalPriceCents += cartItem.DiscountedPriceCents
	}
	return &cartPresenter, nil
}
