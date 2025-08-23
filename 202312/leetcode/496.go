package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l := len(nums1)
	slice496 := make([]int, l)
	for i := 0; i < l; i++ {
		slice496[i] = -1
	}
	for i, v := range nums1 {
		for j := getIndex(nums2, v); j < len(nums2); j++ {
			if nums2[j] > v {
				slice496[i] = nums2[j]
				break
			}
		}
	}
	return slice496
}

func getIndex(arr []int, n int) int {
	index := -1
	for i, v := range arr {
		if n == v {
			index = i
		}
	}
	return index
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
