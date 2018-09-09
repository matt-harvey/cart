package models

import (
	"time"
)

type CartItem struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CartID    int       `json:"cart_id" db:"cart_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	Product   Product   `json:"product" belongs_to:"product"`
	Quantity  uint      `json:"quantity" db:"quantity"`

	// These are total prices for this CartItem, not prices per unit.
	StandardPriceCents   uint `json:"standard_price_cents" db:"standard_price_cents"`
	DiscountedPriceCents uint `json:"discounted_price_cents" db:"discounted_price_cents"`
}

type CartItems []CartItem

// ApplyDiscount assumes Quantity and StandardPriceCents have already been set. It then
// sets DiscountedPriceCents on that basis, by applying the passed Discount to the
// quantity of the CartItem given by discountedQuantity. For example, if buying
// 3 shirts gets you the 4th one at half price, then the discountedQuantity will
// only be 1, not 3. Any excess of discountedQuantity over ci.Quantity is ignored
// in the calculation.
func (ci *CartItem) ApplyDiscount(discount Discount, discountedQuantity uint) {
	if discountedQuantity >= ci.Quantity {
		discountedQuantity = ci.Quantity
	}
	standardPriceCentsPerUnit := ci.StandardPriceCents / ci.Quantity
	discountedPriceCentsPerUnit := discount.Apply(standardPriceCentsPerUnit)
	nonDiscountedQuantity := ci.Quantity - discountedQuantity
	ci.DiscountedPriceCents = discountedQuantity*discountedPriceCentsPerUnit +
		nonDiscountedQuantity*standardPriceCentsPerUnit
}
