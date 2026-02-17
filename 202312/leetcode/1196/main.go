package main

import (
	"fmt"
	"sort"
)

func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	sum := 0
	for i := range weight {
		sum += weight[i]
		if sum > 5000 {
			return i
		}
	}
	return len(weight)
}

func main() {
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800}))
}
