package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	p1, p2 := 0, len(numbers)-1
	for p1 < p2 {
		sum := numbers[p1] + numbers[p2]
		if sum == target {
			return []int{p1 + 1, p2 + 1}
		}
		if sum < target {
			p1++
		} else {
			p2--
		}
	}
	return []int{}
}
