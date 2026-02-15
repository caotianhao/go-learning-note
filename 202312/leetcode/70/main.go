package main

import "fmt"

//func climbStairs(n int) int {
//	dp := make([]int, 3)
//	if n < 3 {
//		return n
//	}
//	dp[1], dp[2] = 1, 2
//	for i := 3; i <= n; i++ {
//		temp := dp[1] + dp[2]
//		dp[1] = dp[2]
//		dp[2] = temp
//	}
//	return dp[2]
//}

//func climbStairs(n int) int {
//	if n < 3 {
//		return n
//	}
//	dp := make([]int, n+1)
//	dp[1], dp[2] = 1, 2
//	for i := 3; i < n+1; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[n]
//}

// 2023.06.15
func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	// 1<=n<=45
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2)) //2
	fmt.Println(climbStairs(3)) //3
}
