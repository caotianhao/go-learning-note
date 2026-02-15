package main

import "fmt"

func applyOperations(nums []int) (ret []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			ret = append(ret, nums[i])
		}
	}
	l1 := len(ret)
	for i := 0; i < l-l1; i++ {
		ret = append(ret, 0)
	}
	return
}

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}
