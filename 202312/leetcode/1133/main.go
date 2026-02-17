package main

import (
	"fmt"
	"sort"
)

func largestUniqueNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := range nums {
		if m[nums[i]] == 1 {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(largestUniqueNumber([]int{8, 9, 9}))
}
