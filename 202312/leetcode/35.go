package main

import (
	"fmt"
	"sort"
)

//func searchInsert(nums []int, target int) int {
//	l := len(nums)
//	low, high := 0, l-1
//	for low <= high {
//		mid := (low + high) / 2
//		if target > nums[mid] {
//			low = mid + 1
//		}
//		if target < nums[mid] {
//			high = mid - 1
//		}
//		if target == nums[mid] {
//			return mid
//		}
//	}
//	return low
//}

//func searchInsert(nums []int, target int) int {
//	return sort.Search(len(nums), func(i int) bool {
//		return nums[i] >= target
//	})
//}

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(idx int) bool {
		return nums[idx] >= target
	})
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
	fmt.Println(searchInsert([]int{1}, -1))
}
