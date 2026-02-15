package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	l, sum := len(nums), 0
	for i := 0; i < l; i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
