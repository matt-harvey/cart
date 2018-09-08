package actions

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

func CreateCartItem(writer http.ResponseWriter, request *http.Request) {
	// TODO Support merging with existing CartItems if Product is same as another
	// already in the Cart.
	cartItem := models.CartItem{}
	err := render.DecodeJSON(request.Body, &cartItem)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	err = db.Conn().Create(&cartItem)
	if err != nil {
		// TODO Handle htis
		Log.Print("DEBUG err: ", err)
		return
	}
	// TODO Respond
}

func DestroyCartItem(writer http.ResponseWriter, request *http.Request) {
	// TODO Support amending quantity of CartItem.
	Log.Print("DEBUG")
	cartItem := models.CartItem{}
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
		return
	}
	err = db.Conn().Find(&cartItem, id)
	if err != nil {
		// TODO Handle htis
		Log.Print("DEBUG err: ", err)
		return
	}
	err = db.Conn().Destroy(&cartItem)
	if err != nil {
		// TODO Handle htis
		Log.Print("DEBUG err: ", err)
		return
	}
	// TODO Respond
}
