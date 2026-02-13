package main

func subArrayRanges(nums []int) int64 {
	sumMax := subArrayMax(nums)
	sumMin := subArrayMin(nums)
	return sumMax - sumMin
}

func subArrayMin(nums []int) int64 {
	nextS := nextSE(nums)
	prevS := prevSE(nums)
	var sum int64 = 0
	for i := 0; i < len(nums); i++ {
		left := int64(i - prevS[i])
		right := int64(nextS[i] - i)
		sum += left * right * int64(nums[i])
	}
	return sum
}

func subArrayMax(nums []int) int64 {
	nextS := nextGE(nums)
	prevS := prevGE(nums)
	var sum int64 = 0
	for i := 0; i < len(nums); i++ {
		left := int64(i - prevS[i])
		right := int64(nextS[i] - i)
		sum += left * right * int64(nums[i])
	}
	return sum
}

func nextGE(nums []int) []int {
	stack := make([]int, 0, len(nums))
	nextG := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			nextG[i] = stack[len(stack)-1]
		} else {
			nextG[i] = len(nums)
		}
		stack = append(stack, i)
	}
	return nextG
}

func prevGE(nums []int) []int {
	stack := make([]int, 0, len(nums))
	prevS := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
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
