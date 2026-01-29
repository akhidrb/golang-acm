package main

func canCompleteCircuit(gas []int, cost []int) int {
	total, m := 0, 0
	in := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		m += gas[i] - cost[i]
		if m < 0 {
			m = 0
			in = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return in
}
