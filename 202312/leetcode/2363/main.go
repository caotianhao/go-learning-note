package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	tmpMap, tmpSlice := map[int]int{}, make([]int, 0)
	ret := make([][]int, 0)
	l1, l2 := len(items1), len(items2)
	for i := 0; i < l1; i++ {
		tmpMap[items1[i][0]] += items1[i][1]
	}
	for i := 0; i < l2; i++ {
		tmpMap[items2[i][0]] += items2[i][1]
	}
	for i := range tmpMap {
		tmpSlice = append(tmpSlice, i)
	}
	sort.Ints(tmpSlice)
	for i := range tmpSlice {
		ret = append(ret, []int{tmpSlice[i], tmpMap[tmpSlice[i]]})
	}
	return ret
}

func main() {
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}))
}
