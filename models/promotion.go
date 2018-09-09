package models

import (
	"time"
)

// Promotion represents a particular discount scheme applicable to certain combinations
// or products.
type Promotion struct {
	ID int `json:"id" db:"id"`

	// Label should be a unique, human-readable label for identifying this promotion
	// within the business.
	Label string `json:"label" db:"label"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// RequiredProductID identifies the Product that a customer is required to purchase
	// in order for this promotion to apply.
	RequiredProductID int `json:"required_product_id" db:"required_product_id"`

	// RequiredProductQuantity identifies the quantity of the required product (identified
	// by RequiredProductID), that a customer needs to purchase in order for this promotion
	// to apply.
	RequiredProductQuantity uint `json:"required_product_quantity" db:"required_product_quantity"`
}
