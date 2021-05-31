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

	log.Println("Loading " + t)

	if t == "Buyer" {

		input, err := ModelBuyer.ReadBody(body)

		if err != nil {
			return err
		}

		req = graphql.NewRequest(ModelBuyer.AddBuyerGQL())
		req.Var("input", ModelBuyer.RemoveDuplicates(input))

	} else if t == "Product" {

		input, err := ModelProduct.ReadBody(body)

		if err != nil {
			return err
		}

		req = graphql.NewRequest(ModelProduct.AddProductGQL())
		req.Var("input", input)

	} else if t == "Transaction" {

		input, err := ModelTransaction.ReadBody(body)

		if err != nil {
			return err
		}

		req = graphql.NewRequest(ModelTransaction.AddTransactionGQL())
		req.Var("input", input)

	} else {
		return errors.New("Not defined")
	}

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	log.Println("Running mutation")
	if err := client.Run(ctx, req, nil); err != nil {
		return err
	}

	log.Println("Successfully inserted")

	return nil
}

func QueryGQL(q string, in map[string]interface{}, output interface{}) error {

	client := GetGraphQL()
	req := graphql.NewRequest(q)
	req.Header.Set("Cache-Control", "no-cache")

	if in != nil {
		for key, val := range in {
			req.Var(key, val)
		}
	}

	ctx := context.Background()

	log.Println("Running query")
	if err := client.Run(ctx, req, &output); err != nil {
		return err
	}

	return nil
}
