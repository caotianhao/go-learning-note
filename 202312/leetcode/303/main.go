package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor303(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (na *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += na.nums[i]
	}
	return sum
}

func main() {
	a := Constructor303([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(a.SumRange(0, 2))
	fmt.Println(a.SumRange(2, 5))
	fmt.Println(a.SumRange(0, 5))
}
