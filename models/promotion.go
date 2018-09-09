package models

import (
	"time"
)

// Promotion represents a particular discount scheme applicable to certain combinations
// of Products.
type Promotion struct {
	ID int `json:"id" db:"id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// RequiredProductID identifies the Product that a customer is required to purchase
	// in order for this promotion to apply.
	RequiredProductID int `json:"required_product_id" db:"required_product_id"`

	// RequiredProductQuantity identifies the quantity of the required product (identified
	// by RequiredProductID), that a customer needs to purchase in order for this promotion
	// to apply.
	RequiredProductQuantity uint `json:"required_product_quantity" db:"required_product_quantity"`

	// DiscountType should be one of the constants of the form ..._DISCOUNT_TYPE.
	DiscountType string `json:"discount_type" db:"discount_type"`

	// DiscountID references an instance of some concrete type satisfying the Discount interface.
	DiscountID int `json:"discount_id" db:"discount_id"`
}

func (p *Promotion) Discount() (Discount, error) {
	return getDiscount(p.DiscountType, p.DiscountID)
}
