package main

import (
	"fmt"
	"sort"
)

func missingNumber268(nums []int) int {
	l, ind := len(nums), -1
	sort.Ints(nums)
	if nums[l-1] == l-1 {
		return l
	}
	for i := 0; i < l; i++ {
		if nums[i] != i {
			ind = i
			break
		}
	}
	return ind
}

func main() {
	fmt.Println(missingNumber268([]int{0, 1, 3}))
}
