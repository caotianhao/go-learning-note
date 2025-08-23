package main

import "fmt"

func missingNumber1704(nums []int) int {
	map1704 := map[int]int{}
	for _, v := range nums {
		map1704[v]++
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := map1704[i]; !ok {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(missingNumber1704([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber1704([]int{3, 0, 1}))
}
