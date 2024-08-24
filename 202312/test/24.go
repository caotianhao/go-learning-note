package main

import "fmt"

func minPath(path [][]int) int {
	if len(path) == 0 || len(path[0]) == 0 {
		return 0
	}
	r, c := len(path), len(path[0])
	dp := make([][]int, r)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, c)
	}
	dp[0][0] = path[0][0]
	for i := 1; i < r; i++ {
		dp[i][0] = dp[i-1][0] + path[i][0]
	}
	for j := 1; j < c; j++ {
		dp[0][j] = dp[0][j-1] + path[0][j]
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + path[i][j]
		}
	}
	return dp[r-1][c-1]
}

func main() {
	fmt.Println(minPath([][]int{{1, 2, 1}, {1, 3, 4}, {5, 7, 8}}))
}
