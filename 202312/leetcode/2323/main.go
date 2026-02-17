package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumTime(jobs, workers []int) int {
	res := -1
	sort.Ints(jobs)
	sort.Ints(workers)
	for i := range jobs {
		t := int(math.Ceil(float64(jobs[i]) / float64(workers[i])))
		if t > res {
			res = t
		}
	}
	return res
}

func main() {
	fmt.Println(minimumTime([]int{3, 18, 15, 9}, []int{6, 5, 1, 3}))
	fmt.Println(minimumTime([]int{5, 2, 4}, []int{1, 7, 5}))
}
