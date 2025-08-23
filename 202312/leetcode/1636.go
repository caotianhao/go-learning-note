package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	cntMap := map[int]int{}
	for i := range nums {
		cntMap[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return (cntMap[nums[i]] < cntMap[nums[j]]) ||
			(cntMap[nums[i]] == cntMap[nums[j]] && nums[i] > nums[j])
	})
	return nums
}

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(frequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
