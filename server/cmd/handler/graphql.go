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

		_, err := ModelTransaction.ReadBody(body)

		if err != nil {
			return err
		}

		return nil

	} else {
		return errors.New("Not defined")
	}

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	var graphqlResponse interface{}
	if err := client.Run(ctx, req, &graphqlResponse); err != nil {
		log.Fatal(err)
	}
	log.Println(graphqlResponse)

	return nil
}
