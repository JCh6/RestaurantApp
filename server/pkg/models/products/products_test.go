package products

import "testing"

func TestAdd(t *testing.T) {
	prods := New()
	prods.Add(Item{})

	if len(prods.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	prods := New()
	prods.Add(Item{})
	results := prods.GetAll()

	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}