package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	pair := 0
	for i := range nums {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pair++
			}
		}
	}
	return pair
}

func main() {
	fmt.Println(numIdenticalPairs([]int{2, 3, 6}))
}
