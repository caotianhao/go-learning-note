package main

import "fmt"

func pivotIndex(nums []int) int {
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
	fmt.Println(pivotIndex([]int{1, 7, 0, 0, 0, 3, 6, 5, 6, 0, 0, 0, 0, 0, 0}))
	fmt.Println(pivotIndex([]int{1}))
	fmt.Println(pivotIndex([]int{0, 1}))
	fmt.Println(pivotIndex([]int{-1, -1, 1, 1, 0, 0}))
}
