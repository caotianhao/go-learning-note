package main

import (
	"fmt"
	"sort"
)

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	fmt.Println(smallestK([]int{1, 2, 3}, 2))
	fmt.Println(smallestK([]int{1, 2, 3}, 0))
	fmt.Println(smallestK([]int{}, 0))
}
