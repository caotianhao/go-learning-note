package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	money, l, candy := 0, len(cost), 1
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	for i := 0; i < l; i++ {
		if candy%3 == 0 {
			candy++
			continue
		}
		money += cost[i]
		candy++
	}
	return money
}

func main() {
	//fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	//fmt.Println(minimumCost([]int{5, 5}))
	//fmt.Println(minimumCost([]int{1, 2, 3}))
	fmt.Println(minimumCost([]int{6, 6, 6, 6, 2, 2}))
}
