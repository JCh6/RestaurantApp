package schema

func Count() string {
	return (`
		query {
			aggregateBuyer {
				count
			}
			aggregateProduct {
				count
			}
			aggregateTransaction {
				count
			}
		}	  
	`)
}
