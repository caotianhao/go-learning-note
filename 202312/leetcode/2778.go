package main

import "fmt"

func sumOfSquares(nums []int) (res int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if n%(i+1) == 0 {
			res += nums[i] * nums[i]
		}
	}
	return
}

func main() {
	fmt.Println(sumOfSquares([]int{1, 2, 3, 4})) // 21
}
