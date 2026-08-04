package main

/* Example:
flagged_vendors = ["acme", "globex", "acme", "initech", "acme", "globex", "umbrella"]
k = 2

Output: ["acme", "globex"]
*/

func flaggedVendors(vendors []string, k int) []string {

	counter := make(map[string]int)
	max_freq := 0
	for _, vendor := range vendors {
		counter[vendor] += 1
		if counter[vendor] > max_freq {
			max_freq = counter[vendor]
		}
	}
	top_k := make([][]string, max_freq+1)
	for vendor, freq := range counter {
		top_k[freq] = append(top_k[freq], vendor)
	}
	result := make([]string, 0)
	for i := len(top_k) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, top_k[i]...)
	}
	return result
}
