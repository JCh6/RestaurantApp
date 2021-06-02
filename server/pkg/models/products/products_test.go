package products

import "testing"

func TestNewProduct(t *testing.T) {
	id := "1x02"
	name := "Flowers"
	price := float32(4141)

	p := New(id, name, price)

	if p.Id != id || p.Name != name || p.Price != price {
		t.Errorf("Product was not created correctly")
	}
}

func TestReadBody(t *testing.T) {
	body := "9e160ac0'Cream of mushroom condensed soup'5020"

	products, err := ReadBody(body)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(products) != 1 {
		t.Errorf("Product was not added")
	}
}
