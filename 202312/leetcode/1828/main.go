package main

import (
	"fmt"
	"math"
)

func countPoints1828(points [][]int, queries [][]int) []int {
	lq := len(queries)
	ret := make([]int, lq)
	for i := 0; i < lq; i++ {
		for _, v := range points {
			if math.Sqrt(math.Pow(float64(v[0]-queries[i][0]), 2)+
				math.Pow(float64(v[1]-queries[i][1]), 2)) <= float64(queries[i][2]) {
				ret[i]++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(countPoints1828([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
		[][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
}
