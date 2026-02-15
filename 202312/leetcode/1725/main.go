package main

import (
	"fmt"
	"sort"
)

func countGoodRectangles(rectangles [][]int) int {
	l, minSlice := len(rectangles), make([]int, 0)
	for i := 0; i < l; i++ {
		minSlice = append(minSlice, minOneSTF(rectangles[i]))
	}
	maxOfMin, cnt := 0, 0
	for i := 0; i < l; i++ {
		if minSlice[i] > maxOfMin {
			maxOfMin = minSlice[i]
		}
	}
	for i := 0; i < l; i++ {
		if minSlice[i] == maxOfMin {
			cnt++
		}
	}
	return cnt
}

func minOneSTF(a []int) int {
	sort.Ints(a)
	return a[0]
}

func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	ret := countGoodRectangles(rectangles)
	fmt.Println(ret)
}
