package main

import "fmt"

func targetIndices(nums []int, target int) []int {
	l, mySlice := len(nums), make([]int, 0)
	bubble2(nums)
	for i := 0; i < l; i++ {
		if nums[i] == target {
			mySlice = append(mySlice, i)
		}
	}
	return mySlice
}

func bubble2(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	nums := []int{1, 2, 5, 2, 3}
	target := 3
	fmt.Println(targetIndices(nums, target))
}
