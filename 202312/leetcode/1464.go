package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return (nums[l-1] - 1) * (nums[l-2] - 1)
}

func main() {
	fmt.Println(maxProduct([]int{1, 2, 3, 6, 9, 663, 25}))
}
