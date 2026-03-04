package main

func findDuplicate(nums []int) int {
	seen := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
		if seen[nums[i]] > 1 {
			return nums[i]
		}
	}
	return 0
}

func findDuplicate3(nums []int) int {
	l, r := 1, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func findDuplicate2(nums []int) int {
	fast, slow := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
