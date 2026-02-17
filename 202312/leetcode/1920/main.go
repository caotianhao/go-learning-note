package main

import "fmt"

func buildArray(nums []int) []int {
	arr := make([]int, 0)
	for i := range nums {
		arr = append(arr, nums[nums[i]])
	}
	return arr
}

func main() {
	arr := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(arr))
}
