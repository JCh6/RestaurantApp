package handler

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/machinebox/graphql"
	"io"
	"log"
	"os"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelProduct "restaurantapp/pkg/models/products"
	"strconv"
	"strings"
)

func GetGraphQL() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL"))
}

func InsertToDgraph(t string, body string) error {

	var req *graphql.Request
	client := GetGraphQL()

	if t == "Buyer" {

		var input []ModelBuyer.Buyer

		if err := json.Unmarshal([]byte(body), &input); err != nil {
			return err
		}

		req = graphql.NewRequest(ModelBuyer.AddBuyerGQL())
		req.Var("input", ModelBuyer.RemoveDuplicates(input))

	} else if t == "Product" {

		var input []ModelProduct.Product
		r := csv.NewReader(strings.NewReader(body))
		r.Comma = '\''

		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			price, err := strconv.ParseFloat(record[2], 32)

			if err != nil {
				return err
			}

			newProduct := ModelProduct.New(record[0], record[1], float32(price))
			input = append(input, *newProduct)
		}

		req = graphql.NewRequest(ModelProduct.AddProductGQL())
		req.Var("input", ModelProduct.RemoveDuplicates(input))

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
