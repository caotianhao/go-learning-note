package main

import (
	"fmt"
	"sort"
)

func findClosestNumber(nums []int) int {
	m2239 := map[int]struct{}{}
	for _, v := range nums {
		m2239[v] = struct{}{}
	}
	for i := range nums {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}
	sort.Ints(nums)
	if _, ok := m2239[nums[0]]; ok {
		return nums[0]
	}
	return -nums[0]
}

func main() {
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8})) //1
	fmt.Println(findClosestNumber([]int{2, -1, 1}))        //1
	fmt.Println(findClosestNumber([]int{2}))               //2
}
