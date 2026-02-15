package main

import (
	"fmt"
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	l1, l2, cnt := len(arr1), len(arr2), 0
	for i := 0; i < l1; i++ {
		yes := true
		for j := 0; j < l2; j++ {
			if int(math.Abs(float64(arr1[i]-arr2[j]))) <= d {
				yes = false
				break
			}
		}
		if yes {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
