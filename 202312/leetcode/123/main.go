package main

import (
	"fmt"
)

const c123 = -int(1e10 + 7)

func maxProfit123(prices []int) (res int) {
	n := len(prices)
	dp := make([][5]int, n)
	// 0 表示什么都没有
	// 1 表示买了一个
	// 2 表示完成一笔交易
	// 3 表示买了第二个
	// 4 表示完成所有交易
	dp[0][1] = -prices[0]
	// 非法状态，防止对结果产生影响所以设为极小的值
	dp[0][2], dp[0][3], dp[0][4] = c123, c123, c123
	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	for i := 0; i < n; i++ {
		t1 := max(dp[i][0], dp[i][2])
		t2 := max(t1, dp[i][4])
		res = max(res, t2)
	}
	return
}

func main() {
	fmt.Println(maxProfit123([]int{7, 1, 5, 3, 6, 4}))
}
