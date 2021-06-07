package transactions

import (
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
	Buyer    string                 `json:"buyer"`
	Ip       string                 `json:"ip"`
	Device   string                 `json:"device"`
	Products []ModelProduct.Product `json:"products"`
}

func New(id string, buyer string, ip string, device string, products []ModelProduct.Product) *Transaction {
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
	keys := make(map[string]bool)
	txnList := strings.Split(body, "#")

	for _, txn := range txnList {

		var products []ModelProduct.Product
		parts := strings.Split(txn, "\x00")

		if len(parts) >= 5 {
			txnId := parts[0]
			buyerId := parts[1]
			ip := parts[2]
			device := parts[3]
			prods := strings.Split(parts[4][1:len(parts[4])-1], ",")

			if _, value := keys[txnId]; !value {
				keys[txnId] = true

				for _, productId := range prods {
					newProduct := ModelProduct.Product{Id: productId}
					products = append(products, newProduct)
				}

				newTransaction := New(txnId, buyerId, ip, device, products)
				input = append(input, *newTransaction)
			}
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

func GetTransactionsByBuyerGQL() string {
	return (`
		query($buyer: String, $first: Int, $offset: Int) {
			queryTransaction(
				filter: { buyer: { eq: $buyer } }
				first: $first
				offset: $offset
			) {
				id
				ip
				device
				productsAggregate {
					count
					priceSum
				}
			}
			aggregateTransaction(filter: { buyer: { eq: $buyer } }) {
				count
			}
		}	  
	`)
}

func GetTransactionByIdGQL() string {
	return (`
		query($id: String!, $ip: String) {
			getTransaction(id: $id) {
				products {
				name
				price
		  		}
			}
			other: queryTransaction(filter: { ip: { eq: $ip } }) {
				buyer
		  		products {
				name
		  		}
			}
	  	}
	`)
}
