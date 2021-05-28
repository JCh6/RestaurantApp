package buyers

import (
	"encoding/json"
)

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func New(id string, name string, age int) *Buyer {
	return &Buyer{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func ReadBody(body string) ([]Buyer, error) {
	var input []Buyer

	if err := json.Unmarshal([]byte(body), &input); err != nil {
		return nil, err
	}

	return input, nil
}

func RemoveDuplicates(list []Buyer) []Buyer {
	keys := make(map[string]bool)
	newList := []Buyer{}

	for _, entry := range list {
		if _, value := keys[entry.Id]; !value {
			keys[entry.Id] = true
			newList = append(newList, entry)
		}
	}

	return newList
}

func AddBuyerGQL() string {
	return (`
		mutation($input: [AddBuyerInput!]!) {
			addBuyer(input: $input) {
				buyer {
					id
				}
			}
		}
	`)
}
