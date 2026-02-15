package main

import "fmt"

func minCostTickets(times, cost []int) int {
	theta := make([]int, times[len(times)-1]+1)
	for idx, t := range times {
		if idx != 0 {
			for i := times[idx-1] + 1; i < t; i++ {
				theta[i] = theta[times[idx-1]]
			}
		}
		xau := min(cost[0]+theta[max(t-1, 0)], cost[1]+theta[max(t-7, 0)])
		theta[t] = min(xau, cost[2]+theta[max(t-30, 0)])
	}
	return theta[times[len(times)-1]]
}

func main() {
	fmt.Println(minCostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 29, 30, 31}, []int{2, 6, 18}))
}
