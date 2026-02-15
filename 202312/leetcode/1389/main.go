package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	l := len(nums)
	target := make([]int, l)
	for i := 0; i < l; i++ {
		target[i] = -1
	}
	for i := 0; i < l; i++ {
		if target[index[i]] == -1 {
			target[index[i]] = nums[i]
		} else {
			for j := i - 1; j >= index[i]; j-- {
				target[j+1] = target[j]
			}
			target[index[i]] = nums[i]
		}
	}
	return target
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	fmt.Println(createTargetArray([]int{1}, []int{0}))
}
