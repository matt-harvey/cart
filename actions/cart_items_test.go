package actions

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

func TestCreateCartItem(t *testing.T) {
	conn := db.Conn()

	// Setup -- TODO Move this into separate function, and remove what's not actually needed
	cart := models.Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}
	beltsProduct := models.Product{Name: "Belts"}
	err = conn.Create(&beltsProduct)
	if err != nil {
		t.Fatal(err)
	}

	json := fmt.Sprintf(`{"cart_id":%d,"product_id":%d,"quantity":%d}`, cart.ID, beltsProduct.ID, 5)
	jsonReader := strings.NewReader(json)
	initialCount, err := db.Conn().Count(models.CartItem{})
	if err != nil {
		t.Fatal("Error counting cart items")
	}

	request, err := http.NewRequest("POST", "/cart_items", jsonReader)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCartItem)
	handler.ServeHTTP(recorder, request)

	newCount, err := conn.Count(models.CartItem{})
	if newCount != initialCount+1 {
		t.Fatal("CartItem not created")
	}

	var quantityResult []int
	err = db.Conn().
		RawQuery("SELECT quantity FROM cart_items ORDER BY created_at DESC LIMIT 1").
		All(&quantityResult)
	if err != nil {
		t.Fatal(err)
	}
	quantity := quantityResult[0]
	expected := 5
	if quantity != expected {
		t.Fatalf("CartItem quantity of %d did not match expected quantity of %d", quantity, expected)
	}
}

func TestDestroyCartItem(t *testing.T) {
	conn := db.Conn()

	// Setup -- TODO Move this into separate function, and remove what's not actually needed
	cart := models.Cart{}
	err := conn.Create(&cart)
	if err != nil {
		t.Fatal(err)
	}
	beltsProduct := models.Product{Name: "Belts"}
	err = conn.Create(&beltsProduct)
	if err != nil {
		t.Fatal(err)
	}
	cartItem := models.CartItem{CartID: cart.ID, ProductID: beltsProduct.ID, Quantity: 5}
	err = conn.Create(&cartItem)
	if err != nil {
		t.Fatal(err)
	}
	initialCount, err := db.Conn().Count(models.CartItem{})
	if err != nil {
		t.Fatal("Error counting cart items")
	}

	url := fmt.Sprintf("/cart_items/%d", cartItem.ID)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DestroyCartItem)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", cartItem.ID))
	request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))
	handler.ServeHTTP(recorder, request)

	newCount, err := db.Conn().Count(models.CartItem{})
	if newCount != initialCount-1 {
		t.Fatal("CartItem not destroyed")
	}
}
