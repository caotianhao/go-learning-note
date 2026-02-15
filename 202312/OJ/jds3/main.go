package main

import "fmt"

func pack(weight, value []int, capacity int) int {
	n := len(weight)
	dp := make([]int, capacity+1)
	selected := make([][]bool, n+1)
	for i := range selected {
		selected[i] = make([]bool, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for w := capacity; w >= weight[i-1]; w-- {
			if dp[w-weight[i-1]]+value[i-1] > dp[w] {
				dp[w] = dp[w-weight[i-1]] + value[i-1]
				selected[i][w] = true
			}
		}
	}
	sl := make([]int, 0)
	w := capacity
	for i := n; i >= 1; i-- {
		if selected[i][w] {
			sl = append(sl, i-1)
			w -= weight[i-1]
		}
	}
	fmt.Println(sl)
	return dp[n]
}

func main() {
	n, capacity := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &capacity)
	ok, okScore, bad, badScore := 0, 0, 0, 0
	weight := make([]int, 0)
	value := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d %d %d %d", &ok, &okScore, &bad, &badScore)
		weight = append(weight, ok, bad)
		value = append(value, okScore, badScore)
	}
	fmt.Printf("%d", pack(weight, value, capacity))
}
