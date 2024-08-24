package main

import "fmt"

func getMax1216(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
		dp[i][i] = dp[i-1][i-1] + matrix[i][i]
		for j := 1; j < i; j++ {
			dp[i][j] = max1216(matrix[i][j]+dp[i-1][j], matrix[i][j]+dp[i-1][j-1])
		}
	}
	maxN := -1
	for _, v := range dp[n-1] {
		if v > maxN {
			maxN = v
		}
	}
	return maxN
}

func max1216(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r, num := 0, 0
	_, _ = fmt.Scan(&r)
	matrix := make([][]int, r)
	for i := 0; i < r; i++ {
		for j := 0; j < i+1; j++ {
			_, _ = fmt.Scan(&num)
			matrix[i] = append(matrix[i], num)
		}
	}
	fmt.Println(getMax1216(matrix))
}
