package actions

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
	cart := models.Cart{}
	db.Conn().Create(&cart)

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
	responseBody := recorder.Body.String()
	createdAt := cart.CreatedAt.Format(time.RFC3339Nano)
	updatedAt := cart.UpdatedAt.Format(time.RFC3339Nano)

	expectedJSON := fmt.Sprintf(
		`{"id":%d,"created_at":"%s","updated_at":"%s"}`,
		cartID,
		createdAt,
		updatedAt)
	if responseBody != expectedJSON {
		t.Fatalf(`JSON response expected to be "%s", was "%s"`, expectedJSON, responseBody)
	}
}
