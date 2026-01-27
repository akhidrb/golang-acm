package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	sum := 0
	candies := make([]int, n)
	for i, _ := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candies[i-1] <= candies[i] {
			candies[i-1] = candies[i] + 1
		}
	}
	for _, val := range candies {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}
