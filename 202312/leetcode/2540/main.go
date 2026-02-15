package main

import "fmt"

func getCommon(nums1, nums2 []int) int {
	k1, k2 := 0, 0
	for k1 < len(nums1) && k2 < len(nums2) {
		if nums1[k1] == nums2[k2] {
			return nums2[k2]
		} else if nums1[k1] > nums2[k2] {
			k2++
		} else {
			k1++
		}
	}
	return -1
}

func main() {
	fmt.Println(getCommon([]int{1}, nil))
}
