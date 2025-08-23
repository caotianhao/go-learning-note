package main

import (
	"fmt"
	"math"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	l, cnt := len(arr), 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
					math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[k]-arr[i])) <= float64(c) {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	fmt.Println(countGoodTriplets(arr, 7, 2, 3))
}
