package main

import "fmt"

func unequalTriplets(nums []int) (cnt int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[k] != nums[j] {
					cnt++
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(unequalTriplets([]int{4, 4, 2, 4, 3}))
}
