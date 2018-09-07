package main

import (
	"net/http"

	"github.com/matt-harvey/cart/config"

	// Ensure the models are initialized form the outset.
	"github.com/matt-harvey/cart/db"
	_ "github.com/matt-harvey/cart/models"
)

func main() {
	router := config.NewRouter()

	db.Conn() // ensure database initialized from the outset

	http.ListenAndServe(config.ApiUrl(), router)
}
