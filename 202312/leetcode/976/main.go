package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	tmp := make([]int, 0)
	for i := range nums {
		tmp = append(tmp, nums[i])
		if len(tmp) == 3 {
			if isTriangle(tmp) {
				return tmp[0] + tmp[1] + tmp[2]
			} else {
				tmp = tmp[1:]
			}
		}
	}
	return 0
}

func isTriangle(tr []int) bool {
	return tr[2]+tr[1] > tr[0]
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{2, 1, 1}))
	fmt.Println(largestPerimeter([]int{1, 2, 3, 4, 5, 9}))
}
