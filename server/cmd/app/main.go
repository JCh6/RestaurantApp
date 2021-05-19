package main

import (
	"restaurantapp/cmd/handler"
	"restaurantapp/pkg/models/products"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main () {
	port := ":3000"
	products := products.New()
	r := chi.NewRouter()

	r.Get("/products", handler.ProductsGet(products))

	http.ListenAndServe(port, r)
}