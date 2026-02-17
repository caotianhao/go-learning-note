package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	ans := make([]int, 0)
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))
	for i := range deck {
		ans = move950(insert950(ans, deck[i]))
	}
	res := append(ans[1:], ans[0])
	return res
}

func move950(arr []int) []int {
	newArr := append([]int{arr[len(arr)-1]}, arr[:len(arr)-1]...)
	return newArr
}

func insert950(arr []int, x int) []int {
	newArr := make([]int, len(arr)+1)
	newArr[0] = x
	for i := 1; i < len(newArr); i++ {
		newArr[i] = arr[i-1]
	}
	return newArr
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}
