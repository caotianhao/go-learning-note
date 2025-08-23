package main

import "fmt"

func arraySign(nums []int) int {
	l, fu := len(nums), 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			fu++
		}
	}
	if fu%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(arraySign([]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}))
}
