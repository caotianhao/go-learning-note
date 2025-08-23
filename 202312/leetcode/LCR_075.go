package main

import (
	"fmt"
	"sort"
)

func relativeSortArray2075(arr1 []int, arr2 []int) []int {
	slice1122 := make([]int, 0)
	for _, v := range arr2 {
		for i := 0; i < howManyNInArr2075(arr1, v); i++ {
			slice1122 = append(slice1122, v)
		}
		arr1 = deleteNFromArr2075(arr1, v)
	}
	sort.Ints(arr1)
	for _, v := range arr1 {
		slice1122 = append(slice1122, v)
	}
	return slice1122
}

func howManyNInArr2075(arr []int, n int) int {
	cnt := 0
	for _, v := range arr {
		if v == n {
			cnt++
		}
	}
	return cnt
}

func deleteNFromArr2075(arr []int, n int) []int {
	slice1122 := make([]int, 0)
	for _, v := range arr {
		if v != n {
			slice1122 = append(slice1122, v)
		}
	}
	return slice1122
}

func main() {
	fmt.Println(relativeSortArray2075([]int{1}, []int{1}))
}
