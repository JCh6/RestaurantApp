package main

import (
	/*"restaurantapp/pkg/models/products"
	"github.com/go-chi/chi/v5"
	"net/http"*/
	//"os"
	//"bytes"
	//"fmt"
	"log"
	//"io/ioutil"
	//"net/http"
	"restaurantapp/cmd/handler"
)

func main() {
	/*port := ":9080"
	products := products.New()
	rr := chi.NewRouter()

	rr.Get("/products", handler.ProductsGet(products))

	http.ListenAndServe(port, rr)*/
	///////////////////
	/*

		urlSetSchema := os.Getenv("SET_SCHEMA_DIR")*/

	/*if handler.PostSchema(urlSetSchema) {
		log.Println("Schema created!")
	}*/

	p := []string{"date", "1231512"}
	handler.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions", p)
	log.Println("Done!")
}
