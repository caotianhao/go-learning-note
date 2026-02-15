package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(nums []int, k int) int {
	l := len(nums)
	if l == 1 {
		return 0
	} else {
		sort.Ints(nums)
		if nums[l-1]-nums[0] <= 2*k {
			return 0
		} else {
			return nums[l-1] - nums[0] - 2*k
		}
	}
}

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
}
