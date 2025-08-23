package main

import "fmt"

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else {
		//for len(nums) != 1 {
		mySlice := make([]int, 0)
		for i := 0; i < len(nums)-1; i++ {
			mySlice = append(mySlice, (nums[i]+nums[i+1])%10)
		}
		nums = mySlice
		return triangularSum(nums)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(triangularSum(arr))
}
