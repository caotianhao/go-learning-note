package main

import (
	"fmt"
	"math"
	"sort"
)

func minStartValue(nums []int) int {
	sum, partSum, l := 0, make([]int, 0), len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
		partSum = append(partSum, sum)
	}
	sort.Ints(partSum)
	return int(math.Max(float64(1-partSum[0]), 1))
}

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, 2}))
	fmt.Println(minStartValue([]int{1, -2, -3}))
}
