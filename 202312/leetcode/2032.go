package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	map1, map2, map3 := map[int]int{}, map[int]int{}, map[int]int{}
	for _, v := range nums1 {
		map1[v] = 1
	}
	for _, v := range nums2 {
		map2[v] = 1
	}
	for _, v := range nums3 {
		map3[v] = 1
	}
	var s1, slice2032 []int
	for i := range map1 {
		s1 = append(s1, i)
	}
	for i := range map2 {
		s1 = append(s1, i)
	}
	for i := range map3 {
		s1 = append(s1, i)
	}
	s1map := map[int]int{}
	for _, v := range s1 {
		s1map[v]++
	}
	for i, v := range s1map {
		if v >= 2 {
			slice2032 = append(slice2032, i)
		}
	}
	return slice2032
}

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
}
