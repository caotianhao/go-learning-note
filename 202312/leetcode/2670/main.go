package main

import "fmt"

func distinctDifferenceArray(nums []int) (diff []int) {
	//l := len(nums)
	for i := 0; i < len(nums); i++ {
		diff = append(diff, set2670(nums[:i+1])-set2670(nums[i+1:]))
	}
	return
}

func set2670(arr []int) int {
	help := make(map[int]struct{})
	for _, v := range arr {
		help[v] = struct{}{}
	}
	return len(help)
}

func main() {
	fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(distinctDifferenceArray([]int{3, 2, 3, 4, 2}))
	fmt.Println(distinctDifferenceArray([]int{3}))
}
