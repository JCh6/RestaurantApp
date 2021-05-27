package products

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

func RemoveDuplicates(list []Product) []Product {
	keys := make(map[string]bool)
	newList := []Product{}

	for _, entry := range list {
		if _, value := keys[entry.Id]; !value {
			keys[entry.Id] = true
			newList = append(newList, entry)
		}
	}

	return newList
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
