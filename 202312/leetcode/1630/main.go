package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))
	temp := make([]int, 0)
	for i := range l {
		temp = append(temp, nums[l[i]:r[i]+1]...)
		res[i] = isArithmetic1630(temp)
		temp = make([]int, 0)
	}
	return res
}

func isArithmetic1630(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return true
	}
	sort.Ints(nums)
	cha := nums[1] - nums[0]
	for i := 1; i < l-1; i++ {
		if nums[i+1]-nums[i] != cha {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7},
		[]int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(checkArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
		[]int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
	//fmt.Println(isArithmetic1630([]int{5, 3, 9, 7}))
	//a := []int{4, 6, 5, 9, 3, 7}
	//fmt.Println(a[2:6])
}
