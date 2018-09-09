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

	// Has the Cart already got a CartItem for this Product?
	productRepresented, err := db.Conn().
		Where("product_id = ? AND cart_id = ?", form.ProductID, cartID).
		Exists(&models.CartItem{})
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	if !productRepresented {
		// Product not yet represented in Cart
		if form.Quantity <= 0 {
			// Can't add non-positive quantity
			// TODO Respond with an error
			Log.Print("DEBUG err: ", err)
			return
		}
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
		return // TODO Respond
	}
	// Product already represented in Cart
	existingItem := models.CartItem{}
	// TODO Shouldn't query again, as we've already done so above.
	err = db.Conn().Where("product_id = ? AND cart_id = ?", form.ProductID, cartID).First(&existingItem)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	var newQuantity uint
	// Stuffing around due to inability to sidestep int/uint conversion issues.
	// TODO What if new form.Quantity is outside int range?
	if form.Quantity < 0 {
		newQuantity = existingItem.Quantity - uint(-form.Quantity)
	} else {
		newQuantity = existingItem.Quantity + uint(form.Quantity)
	}

	if newQuantity < 0 {
		// TODO Handle this -- we can't adjust quantity to less than zero.
		// TODO But what if there are multiple CartItems in Cart for this Product?
		Log.Print("DEBUG err: ", err)
		return
	}
	if newQuantity == 0 {
		// If there's no quantity, delete the CartItem entirely.
		err = db.Conn().Destroy(&existingItem)
	} else {
		existingItem.Quantity = newQuantity
		err = db.Conn().Update(&existingItem)
	}
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
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
