package actions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

func TestCreateCartItem(t *testing.T) {
	conn := db.Conn()

	cart := models.Cart{}
	conn.Create(&cart)

	belts_product := models.Product{Name: "Belts"}
	conn.Create(&belts_product)

	json := fmt.Sprintf(`{"cart_id":%d,"product_id":%d,"quantity":%d}`, cart.ID, belts_product.ID, 5)
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
