package main

import (
	"fmt"
	"reflect"
	"sort"
)

func isMonotonic(nums []int) bool {
	ori := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		ori = append(ori, nums[i])
	}
	sort.Ints(nums)
	if reflect.DeepEqual(nums, ori) {
		return true
	}
	if reflect.DeepEqual(nums, reverse896(ori)) {
		return true
	}
	return false
}

func reverse896(arr []int) []int {
	l := len(arr)
	for i := 0; i < l>>1; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
}
