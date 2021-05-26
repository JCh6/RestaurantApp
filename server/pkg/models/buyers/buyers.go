package buyers

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

func AddBuyer() string {
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
