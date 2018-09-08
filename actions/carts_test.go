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

func TestShowCart(t *testing.T) {
	conn := db.Conn()

	belts_product := models.Product{Name: "Belts", PriceCents: 2000}
	conn.Create(&belts_product)
	trousers_product := models.Product{Name: "Trousers", PriceCents: 7000}
	conn.Create(&trousers_product)

	cart := models.Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}

	trousersItem := models.CartItem{
		CartID:               cart.ID,
		ProductID:            trousers_product.ID,
		Quantity:             3,
		StandardPriceCents:   21000,
		DiscountedPriceCents: 21000,
	}
	err = conn.Create(&trousersItem)
	if err != nil {
		t.Fatal(err)
	}
	beltsItem := models.CartItem{
		CartID:               cart.ID,
		ProductID:            belts_product.ID,
		Quantity:             1,
		StandardPriceCents:   2000,
		DiscountedPriceCents: 1700,
	}
	err = conn.Create(&beltsItem)
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
					"product_name": "Trousers",
					"quantity": 3,
					"standard_price_cents": 21000,
					"discount_cents": 0,
					"discounted_price_cents": 21000
				},
				{
					"product_name": "Belts",
					"quantity": 1,
					"standard_price_cents": 2000,
					"discount_cents": 300,
					"discounted_price_cents": 1700
				}
			]
		}`, cartID)

	if indentedResponseBody != expected {
		t.Fatalf(`JSON response expected to be "%s", was "%s"`, expected, indentedResponseBody)
	}
}
