package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	cnt := 0
	for _, h := range hours {
		if h >= target {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
}
