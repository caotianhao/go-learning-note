package main

import "fmt"

func intersection349(nums1 []int, nums2 []int) []int {
	l1, map1, slice349 := len(nums1), map[int]int{}, make([]int, 0)
	for i := 0; i < l1; i++ {
		map1[nums1[i]]++
	}
	for _, v := range nums2 {
		if _, ok := map1[v]; ok {
			slice349 = append(slice349, v)
		}
	}
	return deleteRedundant(slice349)
}

func deleteRedundant(arr []int) []int {
	l, sliceT, mapT := len(arr), make([]int, 0), map[int]int{}
	for i := 0; i < l; i++ {
		mapT[arr[i]] = arr[i]
	}
	for _, v := range mapT {
		sliceT = append(sliceT, v)
	}
	return sliceT
}

func main() {
	fmt.Println(intersection349([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection349([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection349([]int{0}, []int{0}))
}
