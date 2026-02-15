package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	l, sec := len(points), 0
	for i := 1; i < l; i++ {
		temp := int(math.Abs(float64(points[i][0] - points[i-1][0])))
		temp2 := int(math.Abs(float64(points[i][1] - points[i-1][1])))
		if temp > temp2 {
			sec += temp
		} else {
			sec += temp2
		}
	}
	return sec
}

func main() {
	//fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}
