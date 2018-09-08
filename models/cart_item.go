package models

import (
	"time"
)

type CartItem struct {
	ID                   int       `json:"id" db:"id"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	CartID               int       `json:"cart_id" db:"cart_id"`
	ProductID            int       `json:"product_id" db:"product_id"`
	Quantity             uint      `json:"quantity" db:"quantity"`
	StandardPriceCents   uint      `json:"standard_price_cents" db:"standard_price_cents"`
	DiscountedPriceCents uint      `json:"discounted_price_cents" db:"discounted_price_cents"`
}
