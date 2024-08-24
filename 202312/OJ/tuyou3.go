package main

import (
	"fmt"
	"sort"
)

func pack1(weights, values []int, capacity int) (int, []int) {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	selectItems := make([]int, 0)
	i, w := n, capacity
	for i > 0 && w > 0 {
		if dp[i][w] != dp[i-1][w] {
			selectItems = append(selectItems, i-1)
			w -= weights[i-1]
		}
		i--
	}
	return dp[n][capacity], selectItems
}

func main() {
	n, m, c0, d0 := 0, 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d %d", &n, &m, &c0, &d0)
	a, b, c, d := 0, 0, 0, 0
	list := make([][]int, 0)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
		list = append(list, []int{a, b, c, d})
	}
	// 10 2 1 1
	// 6 3 2 50
	// 8 2 1 10
	sort.Slice(list, func(i, j int) bool {
		return list[i][3] > list[j][3]
	})
	weight := make([]int, 0)
	value := make([]int, 0)
	for _, v := range list {
		for i := 0; i < v[0]/v[1]; i++ {
			weight = append(weight, v[2])
			value = append(value, v[3])
		}
	}
	val, k := pack1(weight, value, n)
	for _, v := range k {
		n -= weight[v]
	}
	for i := 0; i < n; i++ {
		val += d0
	}
	fmt.Println(max(val, n/c0*d0))
}
