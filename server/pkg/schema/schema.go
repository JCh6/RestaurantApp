package schema

func Get() []byte {
	return []byte(`

		type Product {
			id: String! @id
			name: String!
			price: Int!
		}

		type Buyer {
			id: String! @id
			name: String!
			age: Int
		}

		type Transaction {
			id: String! @id
			buyer: Buyer!
			ip: String!
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
