package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	l, tmpSlice, cnt := len(heights), make([]int, 0), 0
	for i := 0; i < l; i++ {
		tmpSlice = append(tmpSlice, heights[i])
	}
	sort.Ints(heights)
	for i := 0; i < l; i++ {
		if tmpSlice[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
