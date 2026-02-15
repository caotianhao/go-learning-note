package main

import (
	"fmt"
	"sort"
)

func missingNumber(arr []int) int {
	sort.Ints(arr)
	l := len(arr)
	cha := (arr[l-1] - arr[0]) / l
	if cha == 0 {
		return arr[0]
	}
	for i := 0; i < l-1; i++ {
		if arr[i+1]-arr[i] != cha {
			return arr[i] + cha
		}
	}
	return -1
}

func main() {
	fmt.Println(missingNumber([]int{15, 14, 12}))
}
