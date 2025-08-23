package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	l := len(citations)
	if citations[l-1] == 0 {
		return 0
	}
	for i := l; i >= 0; i-- {
		if i <= citations[l-i] {
			return i
		}
	}
	return -1
}

func hIndex1(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := n; i > 0; i-- {
		if i <= citations[n-i] {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
	fmt.Println(hIndex([]int{0, 0, 0, 0}))
	fmt.Println(hIndex1([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex1([]int{1, 3, 1}))
	fmt.Println(hIndex1([]int{0, 0, 0, 0}))
}
