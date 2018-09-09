package models

import (
	"math"
	"time"

	"github.com/matt-harvey/cart/db"
)

const PERCENTAGE_DISCOUNT_TYPE = "PercentageDiscount"

// PercentageDiscount represents a discount applicable to a Promotion,
// expressed as a percentage of the standard price of the discounted Product.
type PercentageDiscount struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Amount represents the size of the discount. For example,
	// if Amount is 0.3, then the discount is "30% off".
	Amount float64 `json:"amount" db:"amount"`
}

func init() {
	registerDiscountType(PERCENTAGE_DISCOUNT_TYPE, func(id int) (Discount, error) {
		discount := PercentageDiscount{}
		err := db.Conn().Find(&discount, id)
		if err != nil {
			return nil, err
		}
		return &discount, nil
	})
}

// Apply takes the price for a single unit of a Product, and returns the price that
// unit would have were this PinnedDiscount applied.
func (d *PercentageDiscount) Apply(unitPriceCents uint) uint {
	return unitPriceCents - uint(math.Round(d.Amount*float64(unitPriceCents)))
}
