package handler

import (
	"fmt"
	"net/http"
	"restaurantapp/pkg/models/products"
)

func ProductsGet(products products.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here Products!")
	}
}
