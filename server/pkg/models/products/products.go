package products

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func New(id string, name string, price float32) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func ReadBody(body string) ([]Product, error) {
	var input []Product
	keys := make(map[string]bool)
	r := csv.NewReader(strings.NewReader(body))
	r.Comma = '\''

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		price, err := strconv.ParseFloat(record[2], 32)

		if err != nil {
			return nil, err
		}

		id := record[1]

		newProduct := New(record[0], id, float32(price))

		if _, value := keys[id]; !value {
			keys[id] = true
			input = append(input, *newProduct)
		}
	}

	return input, nil
}

func AddProductGQL() string {
	return (`
		mutation($input: [AddProductInput!]!) {
			addProduct(input: $input) {
		  		product {
					id
		  		}
			}
	  	}
	`)
}
