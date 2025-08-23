package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	//l, ind, flag := len(nums), -1, false
	//if nums[0] == nums[l-1] {
	//	return 1
	//}
	//for i := 0; i < l-1; i++ {
	//	if nums[i] == nums[i+1] {
	//		ind = i
	//		flag = true
	//		break
	//	}
	//}
	//if !flag && sort.IntsAreSorted(nums) {
	//	return l
	//}
	//if nums[ind] == nums[l-1] && sort.IntsAreSorted(nums[:ind+1]) {
	//	return ind + 1
	//}
	//for !canSolveProblem(nums) {
	//	for i := 0; i < l-1; i++ {
	//		if nums[i] == nums[i+1] {
	//			nums = move26(nums, i+1)
	//		}
	//	}
	//}
	//for i := 0; i < l-1; i++ {
	//	if nums[i] > nums[i+1] {
	//		return i + 1
	//	}
	//}
	//return l
	l := len(nums)
	if l == 1 {
		return 1
	}
	up, down := 1, 0
	for up != l {
		if nums[up] == nums[down] {
			up++
			continue
		}
		if nums[up] > nums[down] {
			nums[down+1], nums[up] = nums[up], nums[down+1]
			down++
			up++
		}
	}
	return down + 1
}

//func move26(arr []int, ind int) []int {
//	l := len(arr)
//	if ind == l-1 {
//		return arr
//	}
//	tmp := arr[ind]
//	for i := ind + 1; i < l; i++ {
//		arr[i-1] = arr[i]
//	}
//	arr[l-1] = tmp
//	return arr
//}
//
//func canSolveProblem(arr []int) bool {
//	l, ind := len(arr), -1
//	for i := 0; i < l-1; i++ {
//		if arr[i] >= arr[i+1] {
//			ind = i + 1
//			break
//		}
//	}
//	for i := 0; i < ind; i++ {
//		if arr[i] == arr[i+1] {
//			return false
//		}
//	}
//	return true
//}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3}))
	fmt.Println(removeDuplicates([]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}))
	fmt.Println(removeDuplicates([]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1, 2}))
}
