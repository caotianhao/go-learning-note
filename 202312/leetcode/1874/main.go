package main

import (
	"fmt"
	"sort"
)

func minProductSum(nums1 []int, nums2 []int) int {
	pro := 0
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	for i := range nums1 {
		pro += nums1[i] * nums2[i]
	}
	return pro
}

func main() {
	fmt.Println(minProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}
