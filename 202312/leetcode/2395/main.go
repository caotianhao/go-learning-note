package main

import "fmt"

func findSubarrays(nums []int) bool {
	l := len(nums)
	if l == 2 {
		return false
	}
	for i := 0; i < l-1; i++ {
		tempSum := nums[i] + nums[i+1]
		for j := i + 1; j < l-1; j++ {
			if nums[j]+nums[j+1] == tempSum {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 4, 5}))
	fmt.Println(findSubarrays([]int{0, 0, 0}))
	fmt.Println(findSubarrays([]int{4, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -4}))
}
