package main

import (
	"fmt"
	"sort"
)

//func distinctAverages(nums []int) int {
//	l := len(nums)
//	sort.Ints(nums)
//	myMap := make(map[float64]int, 0)
//	for i := 0; i < l>>1; i++ {
//		myMap[float64(nums[i]+nums[l-1-i])/2]++
//	}
//	return len(myMap)
//}

func distinctAverages(nums []int) int {
	hashMap, l := map[int]struct{}{}, len(nums)
	sort.Ints(nums)
	for i := 0; i < l/2; i++ {
		hashMap[nums[i]+nums[l-i-1]] = struct{}{}
	}
	return len(hashMap)
}

func main() {
	fmt.Println(distinctAverages([]int{4, 1, 4, 0, 3, 5}))
}
