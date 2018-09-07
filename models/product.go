package models

import (
	"time"
)

// Product represents a specific product line purchaseable from the store, for example
// "Blue Acme Jeans". A Product does not represent an individual physical item but rather a line
// of identical items of a particular "make and model".
type Product struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
}
