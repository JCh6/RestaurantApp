package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"restaurantapp/cmd/handler"
	"restaurantapp/pkg/schema"
)

func main() {
	urlSetSchema := os.Getenv("SET_SCHEMA_DIR")

	if handler.PostSchema(urlSetSchema, schema.Get()) {
		log.Println("Schema created!")
		port := os.Getenv("PORT")
		baseUrl := os.Getenv("BASE_ENDPOINT")
		alterDb := os.Getenv("ALTER_DIR")
		r := chi.NewRouter()

		r.Get("/load", handler.GetData(baseUrl, alterDb))
		r.Get("/buyers", handler.GetBuyers())
		r.Get("/transactions", handler.GetTransactions())

		log.Println("Server ready at :" + port)
		log.Fatal(http.ListenAndServe(":"+port, r))
	}
}
