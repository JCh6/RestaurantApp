package products

type Product struct {
	Id    string
	Name  string
	Price float32
}

func New(id string, name string, price float32) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
