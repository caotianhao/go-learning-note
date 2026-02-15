package main

import "fmt"

func maximizeSum(nums []int, k int) (sum int) {
	maxN := 0
	for _, v := range nums {
		if v > maxN {
			maxN = v
		}
	}
	for i := 0; i < k; i++ {
		sum += maxN
		maxN++
	}
	return
}

func main() {
	fmt.Println(maximizeSum([]int{1, 2, 3, 6}, 6))
}
