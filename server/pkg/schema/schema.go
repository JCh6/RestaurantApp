package schema

func Get() []byte {
	return []byte(`

		type City {
			id: ID!
			lat: Float!
			lng: Float!
			secondName: String!
			name: String! @search(by: [exact])
		}

	`)
}