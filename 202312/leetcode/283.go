package main

import "fmt"

//func moveZeroes(nums []int) {
//	zero := 0
//	for _, v := range nums {
//		if v == 0 {
//			zero++
//		}
//	}
//	for i := 0; !isAllZeroInLast(nums, zero); i %= len(nums) - 1 {
//		if nums[i] == 0 && nums[i+1] != 0 {
//			nums[i], nums[i+1] = nums[i+1], nums[i]
//		}
//		i++
//	}
//}

//func isAllZeroInLast(arr []int, n int) bool {
//	l := len(arr)
//	for i := l - n; i < l; i++ {
//		if arr[i] != 0 {
//			return false
//		}
//	}
//	return true
//}

//func moveZeroes(nums []int) {
//	l := len(nums)
//	p, q := 0, 0
//	for i := 0; i < l; i++ {
//		if nums[q] != 0 {
//			nums[p], nums[q] = nums[q], nums[p]
//			p++
//		}
//		q++
//	}
//}

//func moveZeroes(nums []int) {
//	l, zero, no := len(nums), 0, make([]int, 0)
//	for _, v := range nums {
//		if v == 0 {
//			zero++
//		} else {
//			no = append(no, v)
//		}
//	}
//	for i := 0; i < zero; i++ {
//		no = append(no, 0)
//	}
//	for i := 0; i < l; i++ {
//		nums[i] = no[i]
//	}
//}

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func main() {
	nums := []int{0, 1, 0, 0, 3, 12}
	//nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1}
	//nums := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
