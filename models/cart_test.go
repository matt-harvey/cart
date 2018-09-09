package models

import (
	"testing"

	"github.com/matt-harvey/cart/db"
)

func TestCartQualifiesForPromotion(t *testing.T) {
	conn := db.Conn()

	// Qualifying case
	cart := Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}
	jeans := Product{Name: "Jeans"}
	err = conn.Create(&jeans)
	if err != nil {
		t.Fatal(err)
	}
	jeansPromotion := Promotion{
		RequiredProductID:       jeans.ID,
		RequiredProductQuantity: 3,
	}
	jeansItem0 := CartItem{CartID: cart.ID, ProductID: jeans.ID, Quantity: 2}
	jeansItem1 := CartItem{CartID: cart.ID, ProductID: jeans.ID, Quantity: 1}
	err = conn.Create(&jeansItem0)
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Create(&jeansItem1)
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
	err = conn.Create(&hats)
	if err != nil {
		t.Fatal(err)
	}
	hatsItem := CartItem{CartID: cart.ID, ProductID: hats.ID, Quantity: 5}
	err = conn.Create(&hatsItem)
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

func TestCartApplyPricing(t *testing.T) {
	conn := db.Conn()

	// Set up Cart
	cart := Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}

	// Set up Products
	jeans := Product{Name: "Jeans", PriceCents: 9900}
	err = conn.Create(&jeans)
	if err != nil {
		t.Fatal(err)
	}
	socks := Product{Name: "Socks", PriceCents: 500}
	err = conn.Create(&socks)
	if err != nil {
		t.Fatal(err)
	}
	shoes := Product{Name: "Shoes", PriceCents: 20000}
	err = conn.Create(&shoes)
	if err != nil {
		t.Fatal(err)
	}

	// Set up CartItems
	jeansItem := CartItem{CartID: cart.ID, ProductID: jeans.ID, Quantity: 2}
	socksItem := CartItem{CartID: cart.ID, ProductID: socks.ID, Quantity: 3}
	shoesItem := CartItem{CartID: cart.ID, ProductID: shoes.ID, Quantity: 1}
	err = conn.Create(&jeansItem)
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Create(&socksItem)
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Create(&shoesItem)
	if err != nil {
		t.Fatal(err)
	}

	// Set up Discounts
	twentyFivePercentOff := PercentageDiscount{Amount: 0.25}
	err = conn.Create(&twentyFivePercentOff)
	if err != nil {
		t.Fatal(err)
	}

	// Set up qualifying Promotion -- buy two jeans, get 25% off socks
	jeansTwoPromotion := Promotion{
		RequiredProductID:       jeans.ID,
		RequiredProductQuantity: 2,
		DiscountType:            PERCENTAGE_DISCOUNT_TYPE,
		DiscountID:              twentyFivePercentOff.ID}
	err = conn.Create(&jeansTwoPromotion)
	if err != nil {
		t.Fatal(err)
	}

	// Set up non-qualifying Promotion -- buy four socks, get 10% off shoes
	socksFourPromotion := Promotion{
		RequiredProductID:       socks.ID,
		RequiredProductQuantity: 4,
		DiscountType:            PERCENTAGE_DISCOUNT_TYPE,
		DiscountID:              twentyFivePercentOff.ID}
	err = conn.Create(&socksFourPromotion)
	if err != nil {
		t.Fatal(err)
	}

	// Set up PromotionDiscountedProducts
	discountOnSocks := PromotionDiscountedProduct{
		PromotionID: jeansTwoPromotion.ID,
		ProductID:   socks.ID,
	}
	discountOnShoes := PromotionDiscountedProduct{
		PromotionID: socksFourPromotion.ID,
		ProductID:   shoes.ID,
	}
	err = conn.Create(&discountOnSocks)
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Create(&discountOnShoes)
	if err != nil {
		t.Fatal(err)
	}

	// ***ApplyPricing***
	err = cart.ApplyPricing()
	if err != nil {
		t.Fatal(err)
	}

	// Loads cart Items
	if cart.Items == nil {
		t.Fatal("Cart Items not loaded")
	}
	expectedNumCartItems := 3
	if len(cart.Items) != expectedNumCartItems {
		t.Fatalf(
			"Expected %d cart Items to be loaded, but %d were loaded instead",
			expectedNumCartItems,
			len(cart.Items))
	}

	// Test the actual pricing
	item0 := cart.Items[0]
	item1 := cart.Items[1]
	item2 := cart.Items[2]
	if item0.ProductID != jeans.ID {
		t.Fatalf("Expected item0.ProductID to be %d, but was %d", jeans.ID, item0.ProductID)
	}
	if item1.ProductID != socks.ID {
		t.Fatalf("Expected item1.ProductID to be %d, but was %d", socks.ID, item1.ProductID)
	}
	if item2.ProductID != shoes.ID {
		t.Fatalf("Expected item2.ProductID to be %d, but was %d", shoes.ID, item2.ProductID)
	}
	expectedItem0StandardPriceCents := uint(19800)
	if item0.StandardPriceCents != expectedItem0StandardPriceCents {
		t.Fatalf("Expected item0.StandardPriceCents to be %d, but was %d",
			expectedItem0StandardPriceCents, item0.StandardPriceCents)
	}
	expectedItem1StandardPriceCents := uint(1500)
	if item1.StandardPriceCents != expectedItem1StandardPriceCents {
		t.Fatalf("Expected item1.StandardPriceCents to be %d, but was %d",
			expectedItem1StandardPriceCents, item1.StandardPriceCents)
	}
	expectedItem2StandardPriceCents := uint(20000)
	if item2.StandardPriceCents != expectedItem2StandardPriceCents {
		t.Fatalf("Expected item2.StandardPriceCents to be %d, but was %d",
			expectedItem2StandardPriceCents, item2.StandardPriceCents)
	}
	expectedItem0DiscountedPriceCents := uint(19800)
	if item0.DiscountedPriceCents != expectedItem0DiscountedPriceCents {
		t.Fatalf("Expected item0.DiscountedPriceCents to be %d, but was %d",
			expectedItem0DiscountedPriceCents, item0.DiscountedPriceCents)
	}
	expectedItem1DiscountedPriceCents := uint(1125)
	if item1.DiscountedPriceCents != expectedItem1DiscountedPriceCents {
		t.Fatalf("Expected item1.DiscountedPriceCents to be %d, but was %d",
			expectedItem1DiscountedPriceCents, item1.DiscountedPriceCents)
	}
	expectedItem2DiscountedPriceCents := uint(20000)
	if item2.DiscountedPriceCents != expectedItem2DiscountedPriceCents {
		t.Fatalf("Expected item2.DiscountedPriceCents to be %d, but was %d",
			expectedItem2DiscountedPriceCents, item2.DiscountedPriceCents)
	}
}
