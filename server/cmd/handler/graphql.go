package handler

import (
	"context"
	"encoding/json"
	"github.com/machinebox/graphql"
	"log"
	"os"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	"restaurantapp/utils"
)

func GetGraphQL() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL"))
}

func InsertToDgraph(t string, body string) {

	var req *graphql.Request
	client := GetGraphQL()

	if t == "Buyer" {
		var input []ModelBuyer.Buyer

		json.Unmarshal([]byte(body), &input)
		req = graphql.NewRequest(ModelBuyer.AddBuyer())

		req.Var("input", utils.RemoveDuplicates(input))
	}

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	var graphqlResponse interface{}
	if err := client.Run(ctx, req, &graphqlResponse); err != nil {
		log.Fatal(err)
	}
	log.Println(graphqlResponse)

}
