package buyers

import "testing"

func TestNewBuyer(t *testing.T) {
	id := "10x983"
	name := "Hugo Flores"
	age := 24

	b := New(id, name, age)

	if b.Id != id || b.Name != name || b.Age != age {
		t.Errorf("Buyer was not created correctly")
	}
}

func TestReadBody(t *testing.T) {
	body := `[{"id":"10x983","name":"Hugo Flores","age":24}]`

	buyers, err := ReadBody(body)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(buyers) != 1 {
		t.Errorf("Buyer was not added")
	}
}

func TestDuplicates(t *testing.T) {
	body := `[
		{"id":"10x983","name":"Hugo Flores","age":24},
		{"id":"10x983","name":"Maria Flores","age":44}
	]`

	buyers, err := ReadBody(body)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(buyers) != 2 {
		t.Errorf("Buyers were not added")
	}

	buyers = RemoveDuplicates(buyers)

	if len(buyers) != 1 {
		t.Errorf("Did not remove duplicates")
	}
}
