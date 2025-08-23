package main

import "fmt"

func semiOrderedPermutation(nums []int) int {
	left, right, l := 0, 0, len(nums)
	for i, v := range nums {
		if v == 1 {
			left = i
		}
		if v == l {
			right = i
		}
	}
	if left == 0 {
		return l - 1 - right
	}
	if right == l-1 {
		return left
	}
	if left > right {
		return left + (l - 1 - right) - 1
	}
	return left + (l - 1 - right)
}

func main() {
	fmt.Println(semiOrderedPermutation([]int{2, 4, 1, 3}))
}
