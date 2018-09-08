package models

import (
	"time"
)

// StockEntry represents a change in the quantity of a particular Product in inventory.
// Each such change occurs as part of a Transaction.
type StockEntry struct {
	ID            int       `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	TransactionID int       `json:"transaction_id" db:"transaction_id"`
	ProductID     int       `json:"product_id" db:"product_id"`
	Quantity      int       `json:"quantity" db:"quantity"`
}
