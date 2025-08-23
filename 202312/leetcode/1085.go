package main

import "fmt"

func sumOfDigits(nums []int) int {
	minN := 101
	for _, v := range nums {
		if v < minN {
			minN = v
		}
	}
	res := 0
	for minN != 0 {
		res += minN % 10
		minN /= 10
	}
	return (res%2 + 1) % 2
}

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8}))
}
