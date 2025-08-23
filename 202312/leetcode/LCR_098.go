package main

import "fmt"

func uniquePaths2098(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 && j != 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths2098(3, 2))
	fmt.Println(uniquePaths2098(7, 3))
}
