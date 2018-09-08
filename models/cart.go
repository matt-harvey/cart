package models

import (
	"time"
)

type Cart struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Items     CartItems `json:"items" has_many:"cart_items"`
}
