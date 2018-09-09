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

// QualifiesFor returns (true, nil) only if there are Items *persisted* for this
// Cart, such that they cause it to qualify for the passed promotion.
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
