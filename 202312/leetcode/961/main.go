package main

import "fmt"

func repeatedNTimes(nums []int) int {
	repeatedMap := map[int]int{}
	for _, v := range nums {
		repeatedMap[v]++
	}
	for i, v := range repeatedMap {
		if v == len(nums)/2 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}
