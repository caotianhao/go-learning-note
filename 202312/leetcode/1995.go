package main

import (
	"fmt"
)

func countQuadruplets(nums []int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l-3; i++ {
		for j := i + 1; j < l-2; j++ {
			for k := j + 1; k < l-1; k++ {
				for m := k + 1; m < l; m++ {
					if nums[i]+nums[j]+nums[k] == nums[m] {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
	fmt.Println(countQuadruplets([]int{28, 8, 49, 85, 37, 90, 20, 8}))
}
