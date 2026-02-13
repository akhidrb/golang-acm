package main

func sumSubarrayMins(arr []int) int {
	nextS := nextSE(arr)
	prevS := prevSE(arr)
	var MOD int64 = 1_000_000_007
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		left := int64(i - prevS[i])
		right := int64(nextS[i] - i)
		sum = (sum + (left*right*int64(arr[i]))%MOD) % MOD
	}
	return int(sum)
}

func nextSE(nums []int) []int {
	stack := make([]int, 0, len(nums))
	nextS := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			nextS[i] = stack[len(stack)-1]
		} else {
			nextS[i] = len(nums)
		}
		stack = append(stack, i)
	}
	return nextS
}

func prevSE(nums []int) []int {
	stack := make([]int, 0, len(nums))
	prevS := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevS[i] = stack[len(stack)-1]
		} else {
			prevS[i] = -1
		}
		stack = append(stack, i)
	}
	return prevS
}
