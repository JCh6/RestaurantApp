package transactions

import (
	"fmt"
	"log"
	ModelBuyer "restaurantapp/pkg/models/buyers"
	ModelProduct "restaurantapp/pkg/models/products"
	Util "restaurantapp/utils"
	"strings"
)

type Device string

const (
	IOS     Device = "ios"
	ANDROID        = "android"
	WINDOWS        = "windows"
	LINUX          = "linux"
	MAC            = "mac"
)

type Transaction struct {
	Id       string
	buyer    ModelBuyer.Buyer
	ip       string
	device   Device
	products []ModelProduct.Product
}

func ReadBody(body string) ([]Transaction, error) {
	var input []Transaction
	var idx []int
	txnIdLength := 12
	txnList := strings.Split(body, "#")

	for _, txn := range txnList {

		if len(txn) >= txnIdLength {
			txnId := txn[:txnIdLength]
			idx = Util.FindStringIndex(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`, txn)

			if len(idx) > 0 {
				ip := txn[idx[0]:idx[1]]
				buyerId := txn[txnIdLength:idx[0]]
				idx = Util.FindStringIndex(fmt.Sprintf("%s|%s|%s|%s|%s", IOS, ANDROID, WINDOWS, LINUX, MAC), txn)

				if len(idx) > 0 {
					device := txn[idx[0]:idx[1]]
					idx = Util.FindStringIndex(`\(.*\)`, txn)

					if len(idx) > 0 {
						prods := txn[idx[0]+1 : idx[1]-1]
					}

				}
			}
		}

	}

	return input, nil
}
