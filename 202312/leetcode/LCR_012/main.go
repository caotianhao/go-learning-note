package main

import "fmt"

func pivotIndex2012(nums []int) int {
	l, left, right := len(nums), 0, 0
	for i := 0; i < l; i++ {
		left, right = 0, 0
		for j := 0; j < i; j++ {
			left += nums[j]
		}
		for j := i + 1; j < l; j++ {
			right += nums[j]
		}
		if left == right {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex2012([]int{1, 6, 1}))
}
