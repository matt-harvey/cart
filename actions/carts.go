package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/matt-harvey/cart/db"
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
