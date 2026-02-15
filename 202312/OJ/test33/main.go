package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func packHvv(weights, values []int, capacity int) int {
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
	return dp[n][capacity]
}

func main() {
	weights := make([]int, 0)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := strings.Split(input.Text(), ",")
		for _, v := range data {
			tmp, _ := strconv.Atoi(v)
			weights = append(weights, tmp)
		}
	}
	values := make([]int, 0)
	input2 := bufio.NewScanner(os.Stdin)
	for input2.Scan() {
		data := strings.Split(input2.Text(), ",")
		for _, v := range data {
			tmp, _ := strconv.Atoi(v)
			values = append(values, tmp)
		}
	}
	capacity := 0
	_, _ = fmt.Scanf("%d", &capacity)
	fmt.Println(packHvv(weights, values, capacity))
}
