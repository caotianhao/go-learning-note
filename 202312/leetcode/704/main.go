package main

import (
	"fmt"
	"sort"
)

//func search(nums []int, target int) int {
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
//	return -1
//}

//func search(nums []int, target int) int {
//	n := len(nums)
//	l, r, m := 0, n-1, 0
//	for l <= r {
//		m = (l + r) / 2
//		if nums[m] > target {
//			r = m - 1
//		} else if nums[m] < target {
//			l = m + 1
//		} else {
//			return m
//		}
//	}
//	return -1
//}

//func search(nums []int, target int) int {
//	for i, v := range nums {
//		if v == target {
//			return i
//		}
//	}
//	return -1
//}

func search(nums []int, target int) int {
	pos := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if pos < len(nums) && nums[pos] == target {
		return pos
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 0))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 3))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 5))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	//fmt.Println(search([]int{-1}, -3))
	//fmt.Println(search([]int{-1}, -1))
}
