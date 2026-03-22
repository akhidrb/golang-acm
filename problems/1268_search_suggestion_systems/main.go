package main

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	searches := make([][]string, len(searchWord))
	prev := products
	for k := 0; k < len(searchWord); k++ {
		temp := make([]string, 0)
		for _, product := range prev {
			if k < len(product) && product[k] == searchWord[k] {
				temp = append(temp, product)
			}
		}
		prev = temp
		if len(temp) <= 3 {
			searches[k] = temp
		} else {
			searches[k] = temp[:3]
		}
	}

	return searches
}
