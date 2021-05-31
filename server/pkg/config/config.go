package config

type EndPoint struct {
	Name  string
	Route string
	Param string
}

func Get() []EndPoint {
	return []EndPoint{
		EndPoint{
			Name:  "Buyer",
			Route: "/buyers",
			Param: "date",
		},
		EndPoint{
			Name:  "Product",
			Route: "/products",
			Param: "date",
		},
		EndPoint{
			Name:  "Transaction",
			Route: "/transactions",
			Param: "date",
		},
	}
}
