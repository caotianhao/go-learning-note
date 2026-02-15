package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	//初始条件：拆分 2 = 1 + 1, 1 * 1 = 1
	dp[2] = 1
	//从 3 开始
	//思路就是把 n 拆成 i 和 n-i 或 i 和 dp[n-i]
	//前者是分两个数，后者是分三个及以上
	for i := 3; i <= n; i++ {
		//j 从 1 开始，因为 0 乘起来就变成 0
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func main() {
	// [2, 58]
	fmt.Println(integerBreak(10))
}
