package actions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"

	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

func TestCreateCart(t *testing.T) {
	json := "{}"
	jsonReader := strings.NewReader(json)
	initialCount, err := db.Conn().Count(models.Cart{})
	if err != nil {
		t.Fatal("Error counting carts")
	}

	request, err := http.NewRequest("POST", "/carts", jsonReader)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCart)
	handler.ServeHTTP(recorder, request)

	newCount, err := db.Conn().Count(models.Cart{})
	if newCount != initialCount+1 {
		t.Fatal("Cart not created")
	}
}

func TestAdjustCartItemsCart(t *testing.T) {
	conn := db.Conn()

	// Initial creation of Cart
	cart := models.Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}

	// Set up Products
	belts := models.Product{Name: "Belts", PriceCents: 2000}
	conn.Create(&belts)
	trousers := models.Product{Name: "Trousers", PriceCents: 7000}
	conn.Create(&trousers)
	socks := models.Product{Name: "Socks", PriceCents: 800}
	conn.Create(&socks)

	// Set up CartItems
	trousersItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: trousers.ID,
		Quantity:  1,
	}
	err = conn.Create(&trousersItem)
	if err != nil {
		t.Fatal(err)
	}
	beltsItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: belts.ID,
		Quantity:  1,
	}
	err = conn.Create(&beltsItem)
	if err != nil {
		t.Fatal(err)
	}
	// We set up two trousers items, to confirm this scenario is handled correctly.
	otherTrousersItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: trousers.ID,
		Quantity:  2,
	}
	err = conn.Create(&otherTrousersItem)
	if err != nil {
		t.Fatal(err)
	}

	// Scenario: Add a product that's not in cart already
	json := fmt.Sprintf(`{"product_id":%d,"quantity":5}`, socks.ID)
	jsonReader := strings.NewReader(json)
	initialCount, err := db.Conn().Count(models.CartItem{})
	if err != nil {
		t.Fatal("Error counting CartItems")
	}

	url := fmt.Sprintf("/carts/%d/adjust_items", cart.ID)
	Log.Print("DEBUG url: ", url)
	request, err := http.NewRequest("PATCH", url, jsonReader)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AdjustCartItems)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", cart.ID))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))
	handler.ServeHTTP(recorder, request)

	updatedCount, err := db.Conn().Count(models.CartItem{})
	if err != nil {
		t.Fatal("Error counting CartItems")
	}
	expectedCount := initialCount + 1
	if updatedCount != expectedCount {
		t.Fatalf("Count of CartItems, %d, does not match expected count of %d", updatedCount, expectedCount)
	}

}

func TestShowCart(t *testing.T) {
	conn := db.Conn()

	// Initial creation of Cart
	cart := models.Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}

	// Set up Products
	belts := models.Product{Name: "Belts", PriceCents: 2000}
	conn.Create(&belts)
	trousers := models.Product{Name: "Trousers", PriceCents: 7000}
	conn.Create(&trousers)

	// Set up CartItems
	trousersItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: trousers.ID,
		Quantity:  1,
	}
	err = conn.Create(&trousersItem)
	if err != nil {
		t.Fatal(err)
	}
	beltsItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: belts.ID,
		Quantity:  1,
	}
	err = conn.Create(&beltsItem)
	if err != nil {
		t.Fatal(err)
	}
	// We set up two trousers items, to confirm that the quantities in separate
	// Items are merged per Product when we show the Cart as a whole.
	otherTrousersItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: trousers.ID,
		Quantity:  2,
	}
	err = conn.Create(&otherTrousersItem)
	if err != nil {
		t.Fatal(err)
	}

	// Set up Discounts
	fifteenPercentOff := models.PercentageDiscount{Amount: 0.15}
	err = conn.Create(&fifteenPercentOff)
	if err != nil {
		t.Fatal(err)
	}

	// Set up qualifying Promotion -- buy two trousers, get 15% off belts
	trousersPromotion := models.Promotion{
		RequiredProductID:       trousers.ID,
		RequiredProductQuantity: 2,
		DiscountType:            models.PERCENTAGE_DISCOUNT_TYPE,
		DiscountID:              fifteenPercentOff.ID,
	}
	err = conn.Create(&trousersPromotion)
	if err != nil {
		t.Fatal(err)
	}

	// Set up PromotionDiscountedProducts
	discountOnBelts := models.PromotionDiscountedProduct{
		PromotionID: trousersPromotion.ID,
		ProductID:   belts.ID,
	}
	err = conn.Create(&discountOnBelts)
	if err != nil {
		t.Fatal(err)
	}

	cartID := cart.ID
	Log.Print("DEBUG cartID: ", cartID)

	url := fmt.Sprintf("/carts/%d", cartID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ShowCart)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", cartID))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))
	handler.ServeHTTP(recorder, request)
	responseBody := recorder.Body.Bytes()
	var indentedResponseBuffer bytes.Buffer
	err = json.Indent(&indentedResponseBuffer, responseBody, "\t\t", "\t")
	if err != nil {
		t.Fatal(err)
	}
	indentedResponseBody := indentedResponseBuffer.String()

	expected := fmt.Sprintf(`{
			"id": %d,
			"total_price_cents": 22700,
			"items": [
				{
					"product_id": %d,
					"product_name": "Trousers",
					"quantity": 3,
					"standard_price_cents": 21000,
					"discount_cents": 0,
					"discounted_price_cents": 21000
				},
				{
					"product_id": %d,
					"product_name": "Belts",
					"quantity": 1,
					"standard_price_cents": 2000,
					"discount_cents": 300,
					"discounted_price_cents": 1700
				}
			]
		}`, cartID, trousers.ID, belts.ID)

	if indentedResponseBody != expected {
		t.Fatalf(`JSON response expected to be "%s", was "%s"`, expected, indentedResponseBody)
	}
}
