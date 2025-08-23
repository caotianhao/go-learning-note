package main

import "fmt"

func countKDifference(nums []int, k int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if abs2006(nums[i], nums[j], k) {
				cnt++
			}
		}
	}
	return cnt
}

func abs2006(a, b, k int) bool {
	if a-b == k || b-a == k {
		return true
	}
	return false
}

func main() {
	fmt.Println(countKDifference([]int{1, 3}, 6))
}
