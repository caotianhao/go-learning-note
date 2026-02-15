package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	maxN, l := 0, len(nums)
	left, right := 0, l-1
	for left < right {
		maxN = max(maxN, nums[left]+nums[right])
		left++
		right--
	}
	return maxN
}

func main() {
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}
