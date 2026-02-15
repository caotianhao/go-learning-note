package main

import "fmt"

func dietPlanPerformance(calories []int, k, lower, upper int) (score int) {
	sum := make([]int, 1)
	l, first := len(calories), 0
	for i := 0; i < k; i++ {
		first += calories[i]
	}
	sum[0] = first
	ind := k
	for i := 0; i < l-k; i++ {
		sum = append(sum, sum[i]-calories[i]+calories[ind])
		ind++
	}
	for _, v := range sum {
		if v > upper {
			score++
		} else if v < lower {
			score--
		}
	}
	return
}

func main() {
	fmt.Println(dietPlanPerformance([]int{1, 2, 3, 4, 5}, 3, 3, 3))
}
