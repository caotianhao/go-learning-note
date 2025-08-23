package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	allSum, partSum, find, slice1403 := 0, 0, -1, make([]int, 0)
	for _, v := range nums {
		allSum += v
	}
loop:
	for i := 0; i < len(nums); i++ {
		partSum = 0
		for j := 0; j <= i; j++ {
			partSum += nums[j]
			if partSum > allSum-partSum {
				find = i
				break loop
			}
		}
	}
	for i := 0; i <= find; i++ {
		slice1403 = append(slice1403, nums[i])
	}
	return slice1403
}

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))
	fmt.Println(minSubsequence([]int{6}))
}
