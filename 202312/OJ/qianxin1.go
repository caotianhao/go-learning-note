package main

import (
	"fmt"
	"sort"
)

func CheckSquare(planks []int) bool {
	s := 0
	for _, v := range planks {
		s += v
	}
	if s%4 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(planks)))
	edges := [4]int{}
	var square func(int) bool
	square = func(idx int) bool {
		if idx == len(planks) {
			return true
		}
		for i := range edges {
			edges[i] += planks[idx]
			if edges[i] <= s/4 && square(idx+1) {
				return true
			}
			edges[i] -= planks[idx]
		}
		return false
	}
	return square(0)
}

func main() {
	fmt.Println(CheckSquare([]int{1, 1, 2, 2, 2}))
}
