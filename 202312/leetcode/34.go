package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeft34(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	return []int{left, findRight34(nums, target)}
}

//用二分法找到左面第一个出现target的下标
//仔细找区间
func findLeft34(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	for low < high {
		mid := (low + high) >> 1
		if target > nums[mid] {
			low = mid + 1
		}
		if target < nums[mid] {
			high = mid - 1
		}
		if target == nums[mid] {
			high = mid
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}

//用二分法找到右面最后一个出现target的下标
func findRight34(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	for low < high {
		//为啥要加1，这个要理解
		//算是经验，只要打印出low，high，mid看看就知道
		mid := (low + high + 1) >> 1
		if target > nums[mid] {
			low = mid + 1
		}
		if target < nums[mid] {
			high = mid - 1
		}
		if target == nums[mid] {
			low = mid
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}

func main() {
	//                            0  1  2  3   4   5   6   7   8   9
	fmt.Println(searchRange([]int{5, 7, 8, 8, 10, 11, 12, 12, 12, 16}, 12))
	//fmt.Println(findLeft34([]int{5, 7, 8, 8, 10, 11, 12, 12, 12, 16}, 12))
	//fmt.Println(findRight34([]int{5, 7, 8, 8, 10, 11, 12, 12, 12, 16}, 12))
	fmt.Println(searchRange([]int{5, 7, 8, 8, 10}, 8))
}
