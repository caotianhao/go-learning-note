package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	l, minSub, ret := len(arr), 2000001, make([][]int, 0)
	for i := 0; i < l-1; i++ {
		if arr[i+1]-arr[i] <= minSub {
			minSub = arr[i+1] - arr[i]
		}
	}
	for i := 0; i < l-1; i++ {
		if arr[i+1]-arr[i] == minSub {
			ret = append(ret, []int{arr[i], arr[i+1]})
		}
	}
	return ret
}

func main() {
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
	fmt.Println(minimumAbsDifference([]int{40, 11, 26, 27, -20}))
}
