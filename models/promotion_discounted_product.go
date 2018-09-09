package models

import (
	"time"
)

// PromotionDiscountedProduct represents the fact that a given Promotion's Discount is applicable to a
// given Product. In other words, if a Cart qualifies for a Promotion, then the PromotionDiscountedProducts
// linked to this Promotion determine which Products appearing as CartItems in the Cart
// will have their prices modified by the Discount linked to this Promotion.
type PromotionDiscountedProduct struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	ProductID   int       `json:"product_id" db:"product_id"`
	PromotionID int       `json:"promotion_id" db:"promotion_id"`
}
