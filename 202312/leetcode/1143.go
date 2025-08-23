package main

import "fmt"

// 最长公共子序列问题
// 动态规划入门
func longestCommonSubsequence(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)
	// dp[i][j] 代表 s1[:i] 和 s2[:j] 最长公共子序列
	// 若 i==0 相当于 s1 为空串，则不论 s2 为任何串，答案都为 0
	// j==0 同理，所以 dp 第一行和第一列都为 0
	// 由于默认元素为 0，所以不用初始化了
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	// 填充 dp 数组
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			// 如果当前位置两字母相等，则说明长度加 1
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// 如果不等，那就按题意取值，本题为最大，即取最大
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	// 最后返回矩阵右下角的值
	return dp[l1][l2]
}

func longestCommonSubsequence1(text1, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i, byte1 := range text1 {
		for j, byte2 := range text2 {
			if byte1 == byte2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	//fmt.Println(dp)
	return dp[n1][n2]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "acde"))
	fmt.Println(longestCommonSubsequence1("abcde", "acde"))
}
