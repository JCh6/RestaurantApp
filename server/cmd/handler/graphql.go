package handler

import (
	"context"
	"errors"
	"github.com/machinebox/graphql"
	"log"
	"os"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelProduct "restaurantapp/pkg/models/products"
	ModelTransaction "restaurantapp/pkg/models/transactions"
)

func GetGraphQL() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL"))
}

func InsertToDgraph(t string, body string) error {

	var req *graphql.Request
	client := GetGraphQL()

	if t == "Buyer" {

		log.Println("Loading buyers")
		input, err := ModelBuyer.ReadBody(body)

		if err != nil {
			return err
		}

		req = graphql.NewRequest(ModelBuyer.AddBuyerGQL())
		req.Var("input", ModelBuyer.RemoveDuplicates(input))

	} else if t == "Product" {

		log.Println("Loading products")
		input, err := ModelProduct.ReadBody(body)

		if err != nil {
			return err
		}

		req = graphql.NewRequest(ModelProduct.AddProductGQL())
		req.Var("input", input)

	} else if t == "Transaction" {

		log.Println("Loading transactions")
		input, err := ModelTransaction.ReadBody(body)

		if err != nil {
			return err
		}
		log.Printf("Transacciones: %d\n", len(input))
		req = graphql.NewRequest(ModelTransaction.AddTransactionGQL())
		req.Var("input", input)

	} else {
		return errors.New("Not defined")
	}

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	log.Println("Running mutation")
	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully inserted")

	return nil
}
