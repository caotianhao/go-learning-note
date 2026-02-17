package main

import "fmt"

func getConcatenation(nums []int) []int {
	mySlice := make([]int, 0)
	for i := range nums {
		mySlice = append(mySlice, nums[i])
	}
	for i := range nums {
		mySlice = append(mySlice, nums[i])
	}
	return mySlice
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(getConcatenation(arr))
}
