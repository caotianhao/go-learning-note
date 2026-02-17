package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1, nums2 = deleteRedundant2215(nums1), deleteRedundant2215(nums2)
	allMap, deleteSlice := map[int]int{}, make([]int, 0)
	for _, v := range nums1 {
		allMap[v]++
	}
	for _, v := range nums2 {
		allMap[v]++
	}
	for i, v := range allMap {
		if v >= 2 {
			deleteSlice = append(deleteSlice, i)
		}
	}
	for i := range deleteSlice {
		nums1 = deleteNumFromArr2215(nums1, deleteSlice[i])
		nums2 = deleteNumFromArr2215(nums2, deleteSlice[i])
	}
	return [][]int{nums1, nums2}
}

func deleteRedundant2215(arr []int) []int {
	map2215, slice2215 := map[int]int{}, make([]int, 0)
	for _, v := range arr {
		map2215[v]++
	}
	for i := range map2215 {
		map2215[i] = 1
	}
	for i := range map2215 {
		slice2215 = append(slice2215, i)
	}
	return slice2215
}

func deleteNumFromArr2215(arr []int, n int) []int {
	slice2215 := make([]int, 0)
	for _, v := range arr {
		if v == n {
			continue
		}
		slice2215 = append(slice2215, v)
	}
	return slice2215
}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1}))
	fmt.Println(findDifference([]int{0, 0, 0, 0}, []int{0, 0}))
	//fmt.Println(deleteRedundant2215([]int{1, 1, 3, 6, 9, 8, 5, 7, 4}))
}
