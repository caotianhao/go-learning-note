package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	sort.Ints(nums)
	l := len(nums)
	if l&1 == 0 {
		return (float64(nums[l>>1]) + float64(nums[l>>1-1])) / 2
	} else {
		return float64(nums[l>>1])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
}
