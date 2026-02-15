package main

import "fmt"

func longestCommonSubsequence95(text1, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i, b1 := range text1 {
		for j, b2 := range text2 {
			if b1 == b2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[n1][n2]
}

func main() {
	fmt.Println(longestCommonSubsequence95("ad", "add"))
}
