package handler

import (
	"github.com/machinebox/graphql"
	"os"
)

func GetGraphQL() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL"))
}
