package main

import (
	"fmt"
	"math"
)

func minMaxGame(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	slice2293 := make([]int, 0)
	for i := 0; i < l; i += 2 {
		if i%4 == 0 {
			slice2293 = append(slice2293,
				int(math.Min(float64(nums[i]), float64(nums[i+1]))))
		} else {
			slice2293 = append(slice2293,
				int(math.Max(float64(nums[i]), float64(nums[i+1]))))
		}
	}
	return minMaxGame(slice2293)
}

func main() {
	//fmt.Println(minMaxGame([]int{1, 2}))
	fmt.Println(minMaxGame([]int{70, 38, 21, 22}))
}
