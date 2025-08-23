package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	l := len(arr)
	if l == 2 {
		return true
	}
	sort.Ints(arr)
	cha := arr[1] - arr[0]
	for i := 0; i < l-1; i++ {
		if arr[i+1]-arr[i] != cha {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{1, 3, 5}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 888}))
}
