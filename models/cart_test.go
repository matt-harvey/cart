package models

import (
	"testing"

	"github.com/matt-harvey/cart/db"
)

func TestDiscountedPriceCents(t *testing.T) {
	conn := db.Conn()

	cart := Cart{}
	conn.Create(&cart)

	trousers_item := CartItem{CartID: cart.ID, DiscountedPriceCents: 21000}
	conn.Create(&trousers_item)

	belts_item := CartItem{CartID: cart.ID, DiscountedPriceCents: 1700}
	conn.Create(&belts_item)

	expected := 22700
	result, err := cart.DiscountedPriceCents()
	if err != nil {
		t.Fatal(err)
	}

	if expected != result {
		t.Fatalf("Actual result %d did not match expected result %d", result, expected)
	}
}
