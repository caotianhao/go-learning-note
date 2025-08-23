package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minN, maxN := prices[0], -1
	for _, v := range prices {
		if v < minN {
			minN = v
			continue
		}
		if v-minN > maxN {
			maxN = v - minN
		}
	}
	return maxN
}

func maxProfit1(prices []int) int {
	n := len(prices)
	dp := make([]int, n)
	minN := math.MaxInt64
	maxN := 0
	for i, p := range prices {
		if p < minN {
			minN = p
			continue
		}
		dp[i] = p - minN
	}
	for _, v := range dp {
		if v > maxN {
			maxN = v
		}
	}
	return maxN
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit1([]int{7, 6, 5, 4}))
}
