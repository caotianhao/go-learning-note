package main

import (
	"fmt"
	"math"
)

//func minFallingPathSum(matrix [][]int) int {
//	l := len(matrix)
//	dp := make([][]int, l)
//	for i := 0; i < l; i++ {
//		dp[i] = make([]int, l)
//	}
//	for i := 0; i < l; i++ {
//		dp[0][i] = matrix[0][i]
//	}
//	for i := 1; i < l; i++ {
//		dp[i][0] = min931(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
//		dp[i][l-1] = min931(dp[i-1][l-1], dp[i-1][l-2]) + matrix[i][l-1]
//		for j := 1; j < l-1; j++ {
//			dp[i][j] = matrix[i][j] + min931(dp[i-1][j], dp[i-1][j-1], dp[i-1][j+1])
//		}
//	}
//	return min931(dp[l-1]...)
//}
//
//func min931(a ...int) int {
//	min := math.MaxInt64
//	for _, v := range a {
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}

func minFallingPathSum(g [][]int) int {
	n, r := len(g), math.MaxInt64
	for i := 1; i < n; i++ {
		g[i][0] = min(g[i-1][0], g[i-1][1]) + g[i][0]
		g[i][n-1] = min(g[i-1][n-1], g[i-1][n-2]) + g[i][n-1]
		for j := 1; j < n-1; j++ {
			g[i][j] = min(min(g[i-1][j], g[i-1][j-1]), g[i-1][j+1]) + g[i][j]
		}
	}
	for _, v := range g[n-1] {
		if v < r {
			r = v
		}
	}
	return r
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
