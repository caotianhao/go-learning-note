package main

import (
	"fmt"
	"sort"
)

func findLucky(arr []int) int {
	luckyMap, luckySlice := map[int]int{}, make([]int, 0)
	for _, v := range arr {
		luckyMap[v]++
	}
	for i, v := range luckyMap {
		if i == v {
			luckySlice = append(luckySlice, i)
		}
	}
	if len(luckySlice) == 0 {
		return -1
	} else {
		sort.Ints(luckySlice)
		return luckySlice[len(luckySlice)-1]
	}
}

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))
	fmt.Println(findLucky([]int{5}))
	fmt.Println(findLucky([]int{7, 7, 7, 7, 7, 7, 7}))
}
