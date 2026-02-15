package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	mapBin := map[int]int{}
	for _, v := range arr {
		mapBin[v] = howManyOne1356(decToBin1356(v))
	}
	sort.Slice(arr, func(i, j int) bool {
		return (mapBin[arr[i]] < mapBin[arr[j]]) ||
			(mapBin[arr[i]] == mapBin[arr[j]] && arr[i] < arr[j])
	})
	return arr
}

func decToBin1356(n int) []int {
	if n == 0 {
		return []int{0}
	}
	slice1356 := make([]int, 0)
	for n != 0 {
		slice1356 = append(slice1356, n%2)
		n /= 2
	}
	return slice1356
}

func howManyOne1356(arr []int) int {
	cnt := 0
	for _, v := range arr {
		if v == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	//fmt.Println(decToBin1356(11))
	//fmt.Println(decToBin1356(2))
	//fmt.Println(decToBin1356(0))
	fmt.Println(howManyOne1356(decToBin1356(512)))
	//fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	fmt.Println(sortByBits([]int{2, 3, 5, 7, 11, 13, 17, 19}))
	fmt.Println(sortByBits([]int{10, 100, 1000, 10000}))
}
