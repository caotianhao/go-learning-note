package main

import "fmt"

func findIndices(nums []int, indexDiff, valueDiff int) []int {
	res, n := make([]int, 0), len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if abs2903(i-j) >= indexDiff &&
				abs2903(nums[i]-nums[j]) >= valueDiff {
				res = append(res, i, j)
				return res
			}
		}
	}
	return []int{-1, -1}
}

func abs2903(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 4))
}
