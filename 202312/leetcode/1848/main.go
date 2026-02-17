package main

import (
	"fmt"
	"math"
)

func getMinDistance(nums []int, target int, start int) int {
	indexSlice, minN := make([]int, 0), 1000
	for i, v := range nums {
		if v == target {
			indexSlice = append(indexSlice, i)
		}
	}
	for i := range indexSlice {
		tmp := int(math.Abs(float64(start) - float64(indexSlice[i])))
		if tmp < minN {
			minN = tmp
		}
	}
	return minN
}

func main() {
	fmt.Println(getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3))
	fmt.Println(getMinDistance([]int{1}, 1, 0))
}
