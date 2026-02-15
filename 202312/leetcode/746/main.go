package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, 2)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		tmp := min(dp[1]+cost[i-1], dp[0]+cost[i-2])
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         //15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) //6
}
