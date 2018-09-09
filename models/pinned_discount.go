package models

import (
	"time"

	"github.com/matt-harvey/cart/db"
)

const PINNED_DISCOUNT_TYPE = "PinnedDiscount"

// PinnedDiscount represents a discount applicable to a Promotion, whereby the
// discounted Product is priced at a given fixed amount, instead of at its standard
// price.
type PinnedDiscount struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	AmountCents uint      `json:"amount_cents" db:"amount_cents"`
}

func init() {
	registerDiscountType(PINNED_DISCOUNT_TYPE, func(id int) (Discount, error) {
		discount := PinnedDiscount{}
		err := db.Conn().Find(&discount, id)
		if err != nil {
			return nil, err
		}
		return &discount, nil
	})
}

func (d *PinnedDiscount) Apply(priceCents uint) uint {
	if priceCents < d.AmountCents {
		return priceCents
	}
	return d.AmountCents
}
