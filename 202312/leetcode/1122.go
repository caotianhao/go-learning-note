package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	slice1122 := make([]int, 0)
	for _, v := range arr2 {
		for i := 0; i < howManyNInArr1122(arr1, v); i++ {
			slice1122 = append(slice1122, v)
		}
		arr1 = deleteNFromArr(arr1, v)
	}
	sort.Ints(arr1)
	for _, v := range arr1 {
		slice1122 = append(slice1122, v)
	}
	return slice1122
}

func howManyNInArr1122(arr []int, n int) int {
	cnt := 0
	for _, v := range arr {
		if v == n {
			cnt++
		}
	}
	return cnt
}

func deleteNFromArr(arr []int, n int) []int {
	slice1122 := make([]int, 0)
	for _, v := range arr {
		if v != n {
			slice1122 = append(slice1122, v)
		}
	}
	return slice1122
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6}))
}
