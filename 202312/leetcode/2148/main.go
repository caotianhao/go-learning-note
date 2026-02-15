package main

import (
	"fmt"
	"sort"
)

func countElements(nums []int) int {
	sort.Ints(nums)
	l, cnt := len(nums), 0
	for i := 1; i < l-1; i++ {
		if nums[i] > nums[0] && nums[i] < nums[l-1] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countElements([]int{2, 2, 2, 2, 2, 2, 2}))
}
