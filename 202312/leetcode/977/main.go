package main

import (
	"fmt"
)

//func sortedSquares(nums []int) []int {
//	for i, v := range nums {
//		nums[i] = v * v
//	}
//	sort.Ints(nums)
//	return nums
//}

func sortedSquares(nums []int) (res []int) {
	neg, l := -1, len(nums)
	for i := 0; i < l; i++ {
		if nums[i] < 0 {
			neg++
		}
	}
	for i := 0; i < l; i++ {
		nums[i] = nums[i] * nums[i]
	}
	left, right := neg, neg+1
	for left >= 0 && right < l {
		if nums[right] < nums[left] {
			res = append(res, nums[right])
			right++
		} else {
			res = append(res, nums[left])
			left--
		}
	}
	for i := left; i >= 0; i-- {
		res = append(res, nums[i])
	}
	for i := right; i < l; i++ {
		res = append(res, nums[i])
	}
	return
}

func main() {
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
