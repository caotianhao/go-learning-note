package main

import "fmt"

func fib(n int) int {
	//if n == 0 {
	//	return 0
	//}
	//if n == 2 || n == 1 {
	//	return 1
	//} else {
	//	return fib(n-1) + fib(n-2)
	//}
	if n < 2 {
		return n
	}
	dp := make([]int, 2)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		temp := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = temp
	}
	return dp[1]
}

func main() {
	fmt.Println(fib(15))
	fmt.Println(fib(0))
}
