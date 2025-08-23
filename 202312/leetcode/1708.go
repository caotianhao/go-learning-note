package main

import "fmt"

func largestSubarray(nums []int, k int) []int {
	ind, maxB := -1, -1
	for i := 0; i <= len(nums)-k; i++ {
		if nums[i] > maxB {
			maxB = nums[i]
			ind = i
		}
	}
	return nums[ind : ind+k]
}

func main() {
	fmt.Println(largestSubarray([]int{1, 4, 5, 2, 3}, 3))
}
