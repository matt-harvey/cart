package models

import (
	"testing"

	"github.com/matt-harvey/cart/db"
)

func TestCartQualifiesForPromotion(t *testing.T) {

	// Qualifying case
	cart := Cart{}
	err := db.Conn().Create(&cart)
	if err != nil {
		t.Fatal(err)
	}
	jeans := Product{Name: "Jeans"}
	err = db.Conn().Create(&jeans)
	if err != nil {
		t.Fatal(err)
	}
	jeansPromotion := Promotion{
		RequiredProductID:       jeans.ID,
		RequiredProductQuantity: 3,
	}
	jeansItem0 := CartItem{CartID: cart.ID, ProductID: jeans.ID, Quantity: 2}
	jeansItem1 := CartItem{CartID: cart.ID, ProductID: jeans.ID, Quantity: 1}
	err = db.Conn().Create(&jeansItem0)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Conn().Create(&jeansItem1)
	if err != nil {
		t.Fatal(err)
	}
	qualifies, err := cart.QualifiesFor(jeansPromotion)
	if err != nil {
		t.Fatal(err)
	}
	if !qualifies {
		t.Fatal("Expected cart to qualify for jeansPromotion, but it didn't.")
	}

	// Non-qualifying case
	megaJeansPromotion := Promotion{
		RequiredProductID:       jeans.ID,
		RequiredProductQuantity: 4,
	}
	// non-jeans won't help...
	hats := Product{Name: "Hats"}
	err = db.Conn().Create(&hats)
	if err != nil {
		t.Fatal(err)
	}
	hatsItem := CartItem{CartID: cart.ID, ProductID: hats.ID, Quantity: 5}
	err = db.Conn().Create(&hatsItem)
	if err != nil {
		t.Fatal(err)
	}
	qualifies, err = cart.QualifiesFor(megaJeansPromotion)
	if err != nil {
		t.Fatal(err)
	}
	if qualifies {
		t.Fatal("Expected cart not to qualify for megaJeansPromotion, but it did.")
	}
}
