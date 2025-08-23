package main

import "fmt"

func search33(nums []int, target int) int {
	// 题目要求 logn 了那肯定得二分
	n := len(nums)
	left, right, mid := 0, n-1, 0
	// 从 mid 处分为两个数组，其中一个一定是有序的，另一个不一定
	// 根据 mid 对应数字和左右端点的数字大小，即可判断出一定有序的那部分在左还是右
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				// 左部分有序，且 target 在左
				right = mid - 1
			} else {
				// 左部分有序，且 target 在右
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				// 右部分有序，且 target 在右
				left = mid + 1
			} else {
				// 右部分有序，且 target 在左
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search33([]int{4, 5, 6, 7, 0, 1, 2}, 0))    //4
	fmt.Println(search33([]int{4, 5, 6, 7, 0, 1, 2}, 3))    //-1
	fmt.Println(search33([]int{4, 5, 6, 7, 8, 0, 1, 2}, 2)) //7
}
