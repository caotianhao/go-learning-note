package main

import (
	"fmt"
	"sort"
)

func minimumOperations(nums []int) int {
	l := len(nums)
	if isAllZero2357(nums) {
		return 0
	}
	if l == 1 {
		return 1
	}
	cnt := 0
	sort.Ints(nums)
	for !isAllZero2357(nums) {
		tmp := 0
		for i := 0; i < l; i++ {
			if nums[i] != 0 {
				tmp = nums[i]
				for j := i; j < l; j++ {
					nums[j] -= tmp
				}
				cnt++
			}
		}
	}
	return cnt
}

func isAllZero2357(nums []int) bool {
	for i := range nums {
		if nums[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
	fmt.Println(minimumOperations([]int{1, 2, 3, 5}))
	fmt.Println(minimumOperations([]int{7, 5, 0, 3, 5, 6, 12, 1}))
	fmt.Println(minimumOperations([]int{1, 2, 3, 5}))
	//fmt.Println(isAllZero2357([]int{0, 0, 0, 0, 0, 0}))
	//fmt.Println(isAllZero2357([]int{0}))
}
