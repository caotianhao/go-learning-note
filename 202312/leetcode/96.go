package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+2)
	dp[0], dp[1], dp[2] = 1, 1, 2
	if n < 3 {
		return dp[n]
	}
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(19))
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
}
