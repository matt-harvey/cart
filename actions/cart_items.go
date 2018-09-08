package actions

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/matt-harvey/cart/db"
	"github.com/matt-harvey/cart/models"
)

func CreateCartItem(writer http.ResponseWriter, request *http.Request) {
	cartItem := models.CartItem{}
	err := render.DecodeJSON(request.Body, &cartItem)
	if err != nil {
		// TODO Handle this
		Log.Print("DEBUG err: ", err)
	}
	err = db.Conn().Create(&cartItem)
	if err != nil {
		// TODO Handle htis
		Log.Print("DEBUG err: ", err)
	}
	// TODO Respond
}
