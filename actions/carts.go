package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/forms"
	"github.com/matt-harvey/cart/models"
	"github.com/matt-harvey/cart/presenters"
)

// "/cart"
func CreateCart(writer http.ResponseWriter, request *http.Request) {
	// TODO What about validation?
	cart := models.Cart{}
	err := render.DecodeJSON(request.Body, &cart)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	err = db.Conn().Create(&cart)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	// TODO Respond
}

// "/cart/{id}/adjust_items"
func AdjustCartItems(writer http.ResponseWriter, request *http.Request) {
	Log.Print("DEBUG")
	Log.Print("DEBUG request: ", request)
	cartID, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	Log.Print("DEBUG")
	cart := models.Cart{}
	err = db.Conn().Find(&cart, cartID)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	form := forms.CartAdjustItems{}
	err = render.DecodeJSON(request.Body, &form)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	Log.Print("DEBUG")
	if form.Quantity >= 0 {
		cartItem := models.CartItem{
			CartID:    cartID,
			ProductID: form.ProductID,
			Quantity:  uint(form.Quantity)}
		err = db.Conn().Create(&cartItem)
		if err != nil {
			// TODO Handle this
			Log.Print("DEBUG err: ", err)
			return
		}
	}
	// FIXME
}

// "/cart/{id}"
func ShowCart(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	cartPresenter, err := presenters.NewCartPresenter(id)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	json, err := json.Marshal(cartPresenter)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	Log.Print("DEBUG json: ", string(json))
	writer.Write([]byte(json))
}
