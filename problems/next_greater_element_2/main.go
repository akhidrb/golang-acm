package main

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0, len(nums))
	indexMap := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		indexMap[i] = nums[i]
		top := -1_000_000_000
		if len(stack) > 0 {
			top = stack[len(stack)-1]
			for top <= nums[i] {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				} else {
					top = -1_000_000_000
					break
				}
			}
		}
		stack = append(stack, nums[i])
		nums[i] = top
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == -1_000_000_000 {
			if len(stack) > 0 {
				top := -1_000_000_000
				for top <= indexMap[i] {
					if len(stack) > 0 {
						top = stack[len(stack)-1]
						if top > indexMap[i] {
							nums[i] = top
						} else {
							stack = stack[:len(stack)-1]
						}
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == -1_000_000_000 {
			nums[i] = -1
		}
	}

	return nums
}
