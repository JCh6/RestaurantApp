package handler

import (
	"fmt"
	"restaurantapp/pkg/models/products"
	"net/http"
)

func ProductsGet(products products.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here Products!")
	}
}