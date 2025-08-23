package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) (max int) {
	first := make([]int, 0)
	for _, v := range points {
		first = append(first, v[0])
	}
	sort.Ints(first)
	for i := 0; i < len(first)-1; i++ {
		tmp := first[i+1] - first[i]
		if tmp > max {
			max = tmp
		}
	}
	return
}

func main() {
	fmt.Println(maxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}})) //1
}
