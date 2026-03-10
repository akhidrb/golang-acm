package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := dict[nums[i]]; ok {
			if i != val {
				return []int{i, val}
			}
		}
		dict[target-nums[i]] = i
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
