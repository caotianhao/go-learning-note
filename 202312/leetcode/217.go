package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hashMap := map[int]int{}
	for i, v := range nums {
		if _, ok := hashMap[v]; ok {
			return true
		}
		hashMap[v] = i
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12}))
}
