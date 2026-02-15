package main

import (
	"fmt"
	"math"
	"sort"
)

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	res := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		t := nums[i] - nums[i-1]
		if t < res {
			res = t
		}
	}
	return res
}

func main() {
	fmt.Println(findValueOfPartition([]int{1, 3, 2, 4}))
}
