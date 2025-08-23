package main

import (
	"fmt"
	"sort"
)

func searchInsert068(nums []int, target int) int {
	return sort.Search(len(nums), func(idx int) bool {
		return nums[idx] >= target
	})
}

func main() {
	fmt.Println(searchInsert068([]int{1, 3, 5, 7}, 5))
}
