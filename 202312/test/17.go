package main

import "fmt"

func pack17(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if weights[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[n][capacity]
}

func main() {
	capacity, book := 0, 0
	_, _ = fmt.Scanf("%d %d", &capacity, &book)
	w, v := 0, 0
	weights, values := make([]int, 0), make([]int, 0)
	for i := 0; i < book; i++ {
		_, _ = fmt.Scanf("%d %d", &w, &v)
		weights = append(weights, w)
		values = append(values, v)
	}
	fmt.Println(pack17(weights, values, capacity))
}
