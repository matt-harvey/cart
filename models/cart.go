package models

import (
	"time"

	"github.com/matt-harvey/cart/db"
)

type Cart struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Items     CartItems `json:"items" has_many:"cart_items"`
}

func (c *Cart) DiscountedPriceCents() (int, error) {
	// TODO Consider doing this in Go not SQL: if we're always loading items into memory
	// anyway, we might as well iterate over items to calculate.
	var result []int
	err := db.Conn().
		RawQuery("SELECT SUM(discounted_price_cents) FROM cart_items WHERE cart_id = ?", c.ID).
		All(&result)
	return result[0], err
}
