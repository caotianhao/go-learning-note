package main

import "fmt"

func leftRigthDifference(nums []int) []int {
	l := len(nums)
	left, right, res := make([]int, l), make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			left[i] += nums[j]
		}
		for k := l - 1; k > i; k-- {
			right[i] += nums[k]
		}
	}
	for i := 0; i < l; i++ {
		res[i] = abs2574(left[i] - right[i])
	}
	return res
}

func abs2574(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(leftRigthDifference([]int{10, 4, 8, 3}))
}
