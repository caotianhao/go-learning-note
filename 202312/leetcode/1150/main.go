package main

import "fmt"

func isMajorityElement(nums []int, target int) bool {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	return m[target] > len(nums)/2
}

func main() {
	fmt.Println(isMajorityElement([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5))
}
