package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, len(nums))
	for i := range nums {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt++
			}
		}
		arr[i] = cnt
	}
	return arr
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{1, 2, 36, 2}))
}
