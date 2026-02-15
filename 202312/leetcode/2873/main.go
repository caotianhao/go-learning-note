package main

import (
	"fmt"
)

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	maxN := int64(-0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				t := (nums[i] - nums[j]) * nums[k]
				if int64(t) > maxN {
					maxN = int64(t)
				}
			}
		}
	}
	return maxN
}

func main() {
	fmt.Println(maximumTripletValue([]int{1, 2, 3}))
}
