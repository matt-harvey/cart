package config

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/matt-harvey/cart/actions"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/carts", func(router chi.Router) {
		router.Post("/", actions.CreateCart)
		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", actions.ShowCart)
		})
	})
	router.Route("/cart_items", func(router chi.Router) {
		router.Post("/", actions.CreateCartItem)
	})
	return router
}
