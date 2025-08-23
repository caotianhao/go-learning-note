package main

import (
	"fmt"
	"sort"
)

func arraysIntersection(arr1, arr2, arr3 []int) (arr []int) {
	m := map[int]int{}
	for _, v := range arr1 {
		m[v]++
	}
	for _, v := range arr2 {
		m[v]++
	}
	for _, v := range arr3 {
		m[v]++
	}
	for i, v := range m {
		if v == 3 {
			arr = append(arr, i)
		}
	}
	sort.Ints(arr)
	return
}

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
}
