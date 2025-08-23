package main

import (
	"fmt"
)

func pivotArray(nums []int, pivot int) (res []int) {
	less, equal, great := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if v < pivot {
			less = append(less, v)
		} else if v > pivot {
			great = append(great, v)
		} else {
			equal = append(equal, v)
		}
	}
	res = append(res, less...)
	res = append(res, equal...)
	res = append(res, great...)
	return
}

func main() {
	fmt.Println(pivotArray([]int{1, 2, 3, 4, 5, 6}, 5))
}
