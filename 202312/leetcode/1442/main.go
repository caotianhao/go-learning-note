package main

import "fmt"

func countTriplets(arr []int) (cnt int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for k := i + 1; k < l; k++ {
			if verify(i, k, arr) == 0 {
				cnt += k - i
			}
		}
	}
	return
}

func verify(i, j int, arr []int) int {
	a := 0
	for k := i; k <= j; k++ {
		a ^= arr[k]
	}
	return a
}

func main() {
	arr := []int{2, 3, 1, 6, 7}
	fmt.Println(countTriplets(arr))
}
