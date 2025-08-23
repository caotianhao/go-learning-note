package main

import (
	"fmt"
	"sort"
)

func mergeArrays(nums1 [][]int, nums2 [][]int) (ret [][]int) {
	myMap, row := map[int]int{}, make([]int, 0)
	l1, l2 := len(nums1), len(nums2)
	for i := 0; i < l1; i++ {
		myMap[nums1[i][0]] += nums1[i][1]
	}
	for i := 0; i < l2; i++ {
		myMap[nums2[i][0]] += nums2[i][1]
	}
	for i := range myMap {
		row = append(row, i)
	}
	sort.Ints(row)
	for _, v := range row {
		ret0 := make([]int, 0)
		ret0 = append(ret0, v, myMap[v])
		ret = append(ret, ret0)
	}
	return
}

func main() {
	fmt.Println(mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}))
	fmt.Println(mergeArrays([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))
}
