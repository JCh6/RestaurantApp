package main

import (
	/*"restaurantapp/pkg/models/products"
	"github.com/go-chi/chi/v5"
	"net/http"*/
	"os"
	//"bytes"
	//"fmt"
	"log"
	//"io/ioutil"
	//"net/http"
	"restaurantapp/cmd/handler"
)

func main () {
	/*port := ":9080"
	products := products.New()
	rr := chi.NewRouter()

	rr.Get("/products", handler.ProductsGet(products))

	http.ListenAndServe(port, rr)*/
	///////////////////

	//fmt.Println(os.Getenv("KAKA"))
	
	var urlSetSchema string = ""

	urlSetSchema = os.Getenv("SET_SCHEMA_DIR")

	if handler.PostSchema(urlSetSchema) {
		log.Println("Schema created!")
	}

}