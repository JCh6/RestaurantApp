package schema

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data   Status   `json:"data"`
	Errors []Status `json:"errors"`
}

func Get() []byte {
	return []byte(`

		type Product {
			id: String! @id
			name: String!
			price: Float!
		}

		type Buyer {
			id: String! @id
			name: String!
			age: Int
		}

		type Transaction {
			id: String! @id
			buyer: String! @search(by: [hash])
			ip: String! @search(by: [exact])
			device: Device
			products: [Product!]!
		}

		enum Device {
			ios
			android
			windows
			linux
			mac
		}

	`)
}

func DeleteData() []byte {
	return []byte(`

		{ "drop_op": "DATA" }

	`)
}
