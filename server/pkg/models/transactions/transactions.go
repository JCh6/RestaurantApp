package transactions

import (
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelProduct "restaurantapp/pkg/models/products"
	"strings"
)

const (
	IOS     = "ios"
	ANDROID = "android"
	WINDOWS = "windows"
	LINUX   = "linux"
	MAC     = "mac"
)

type Transaction struct {
	Id       string                 `json:"id"`
	Buyer    ModelBuyer.Buyer       `json:"buyer"`
	Ip       string                 `json:"ip"`
	Device   string                 `json:"device"`
	Products []ModelProduct.Product `json:"products"`
}

func New(id string, buyer ModelBuyer.Buyer, ip string, device string, products []ModelProduct.Product) *Transaction {
	return &Transaction{
		Id:       id,
		Buyer:    buyer,
		Ip:       ip,
		Device:   device,
		Products: products,
	}
}

func ReadBody(body string) ([]Transaction, error) {
	var input []Transaction
	var products []ModelProduct.Product
	txnList := strings.Split(body, "#")

	for _, txn := range txnList {

		parts := strings.Split(txn, "\x00")

		if len(parts) >= 5 {
			txnId := parts[0]
			buyerId := parts[1]
			ip := parts[2]
			device := parts[3]
			prods := strings.Split(parts[4][1:len(parts[4])-1], ",")

			for _, product := range prods {
				newProduct := ModelProduct.Product{Id: product}
				products = append(products, newProduct)
			}

			newBuyer := ModelBuyer.Buyer{Id: buyerId}
			newTransaction := New(txnId, newBuyer, ip, device, products)
			input = append(input, *newTransaction)
		}

	}

	return input, nil
}

func AddTransactionGQL() string {
	return (`
		mutation($input: [AddTransactionInput!]!) {
			addTransaction(input: $input) {
		  		transaction {
					id
		  		}
			}
	  	}
	`)
}
