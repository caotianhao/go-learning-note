package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	n, res := len(nums), math.MinInt64
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-2, 1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
