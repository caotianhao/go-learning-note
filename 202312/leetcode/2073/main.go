package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	l, cnt, i := len(tickets), 0, 0
	for tickets[k] != 0 {
		if tickets[i%l] > 0 {
			tickets[i%l]--
			cnt++
		}
		i++
	}
	return cnt
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))
}
