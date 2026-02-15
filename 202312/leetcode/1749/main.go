package main

import (
	"fmt"
	"math"
)

func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	partSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		partSum[i+1] += partSum[i] + nums[i]
	}
	maxN, minN := math.MinInt64, math.MaxInt64
	for _, v := range partSum {
		if v > maxN {
			maxN = v
		}
		if v < minN {
			minN = v
		}
	}
	return maxN - minN
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))
	fmt.Println(maxAbsoluteSum([]int{1}))
}
