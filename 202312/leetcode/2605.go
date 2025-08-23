package main

import "fmt"

func minNumber(nums1 []int, nums2 []int) int {
	hm := map[int]int{}
	min1, min2 := 10, 10
	for _, v := range nums1 {
		hm[v]++
		if v < min1 {
			min1 = v
		}
	}
	for _, v := range nums2 {
		hm[v]++
		if v < min2 {
			min2 = v
		}
	}
	for i := 1; i <= 9; i++ {
		if hm[i] == 2 {
			return i
		}
	}
	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}

func main() {
	fmt.Println(minNumber([]int{1, 2, 3}, []int{2, 3, 6, 9}))
}
