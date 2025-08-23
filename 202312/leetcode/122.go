package main

import (
	"fmt"
)

func maxProfit122(prices []int) int {
	n := len(prices)
	// dp[i][0] 表示当天没有股票的最大利润
	// dp[i][1] 表示当天有股票的最大利润
	dp := make([][2]int, n)
	// 初始状态
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 手里没有股票，可能是卖出去了，也可能是没操作
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		// 手里有股票，可能是今天买的，也可能是没操作
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

func maxProfit12201(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func main() {
	fmt.Println(maxProfit122([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit12201([]int{7, 1, 5, 3, 6, 4}))
}
