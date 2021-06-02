package transactions

import (
	ModelProduct "restaurantapp/pkg/models/products"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	id := "012398x90132"
	buyer := "10x983"
	ip := "127.0.0.1"
	device := "ios"
	p := ModelProduct.New("9431x32", "Apple Computer", float32(2554323))
	products := []ModelProduct.Product{*p}

	txn := New(id, buyer, ip, device, products)

	if txn.Id != id || txn.Buyer != buyer || txn.Ip != ip || txn.Device != device || len(txn.Products) != 1 {
		t.Errorf("Buyer was not created correctly")
	}
}

func TestReadBody(t *testing.T) {
	body := "#000060b6ca00\x00490d6704\x00157.62.23.254\x00mac\x00(9e160ac0)\x00\x00#000060b6ca01\x00dfce24c5\x0025.244.243.23\x00linux\x00(e432ddca,2874756c)\x00\x00"

	transactions, err := ReadBody(body)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(transactions) != 2 {
		t.Errorf("Transactions were not added")
	}
}
