package main

import (
	"fmt"
	"sort"
)

func diagonalPrime(nums [][]int) int {
	l, arr := len(nums), make([]int, 0)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j || i+j == l-1 {
				arr = append(arr, nums[i][j])
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i := range arr {
		if isPrime2614(arr[i]) {
			return arr[i]
		}
	}
	return 0
}

func isPrime2614(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}))
	fmt.Println(isPrime2614(2))
}
