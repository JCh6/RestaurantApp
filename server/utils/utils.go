package utils

import (
	ModelBuyer "restaurantapp/pkg/models/buyers"
)

func RemoveDuplicates(list []ModelBuyer.Buyer) []ModelBuyer.Buyer {
	keys := make(map[string]bool)
	newList := []ModelBuyer.Buyer{}

	for _, entry := range list {
		if _, value := keys[entry.Id]; !value {
			keys[entry.Id] = true
			newList = append(newList, entry)
		}
	}

	return newList
}
