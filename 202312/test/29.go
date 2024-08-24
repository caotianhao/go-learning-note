package main

import "fmt"

func maxDivisibleBy3(s string) int {
	n := len(s)
	dp := make([]int, 3)

	for i := 0; i < n; i++ {
		digit := int(s[i] - '0')

		// 更新 dp 数组，分别考虑当前数字加到 dp 数组的哪一个位置
		newDP := make([]int, 3)
		for j := 0; j < 3; j++ {
			newDP[(j+digit)%3] = max(dp[j]+1, newDP[(j+digit)%3])
			newDP[j] = max(dp[j], newDP[j])
		}

		dp = newDP
	}

	return dp[0]
}

func main() {
	s := "1123" // 替换为你的字符串
	result := maxDivisibleBy3(s)
	fmt.Println("最多可以获得的3的倍数个数：", result)
}
