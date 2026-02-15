package main

import (
	"fmt"
	"sort"
)

func intersection(nums [][]int) []int {
	l, tmpSlice, tmpMap := len(nums), make([]int, 0), map[int]int{}
	for j := 0; j < len(nums[0]); j++ {
		tmpSlice = append(tmpSlice, nums[0][j])
	}
	for i := 0; i < l-1; i++ {
		for j := 0; j < len(nums[i+1]); j++ {
			tmpSlice = append(tmpSlice, nums[i+1][j])
		}
		tmpMap = map[int]int{}
		for j := 0; j < len(tmpSlice); j++ {
			tmpMap[tmpSlice[j]]++
		}
		tmpSlice = make([]int, 0)
		for j, v := range tmpMap {
			if v == 2 {
				tmpSlice = append(tmpSlice, j)
			}
		}
	}
	sort.Ints(tmpSlice)
	return tmpSlice
}

func main() {
	fmt.Println(intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
	fmt.Println(intersection([][]int{{1, 2, 3}, {4, 5, 6}}))
}
