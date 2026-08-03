package main

import (
	"cmp"
	"slices"
)

/* Example:
flagged_vendors = ["acme", "globex", "acme", "initech", "acme", "globex", "umbrella"]
k = 2

Output: ["acme", "globex"]
*/

type VendorFlag struct {
	vendor string
	freq   int
}

func flaggedVendors(vendors []string, k int) []string {
	top_k := make([]string, 0)
	counter := make(map[string]int)
	for _, vendor := range vendors {
		counter[vendor] += 1
	}
	flags := make([]VendorFlag, 0)
	for key, value := range counter {
		flags = append(flags, VendorFlag{vendor: key, freq: value})
	}
	slices.SortFunc(flags, func(a, b VendorFlag) int {
		return cmp.Compare(b.freq, a.freq)
	})

	for i := 0; i < k; i++ {
		top_k = append(top_k, flags[i].vendor)
	}
	return top_k
}
