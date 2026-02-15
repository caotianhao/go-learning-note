package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	l, all := len(salary), 0
	sort.Ints(salary)
	salary = salary[1 : l-1]
	for _, v := range salary {
		all += v
	}
	return float64(all) / float64(l-2)
}

func main() {
	fmt.Println(average([]int{1000, 2000, 3000}))
}
