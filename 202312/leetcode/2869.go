package main

import "fmt"

func minOperations2869(nums []int, k int) int {
	b, n := make([]bool, k), len(nums)
	for i := 0; i < n; i++ {
		num := nums[n-1-i]
		if num <= k {
			b[num-1] = true
		}
		var f = true
		for _, v := range b {
			f = f && v
		}
		if f {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(minOperations2869([]int{3, 1, 5, 4, 2}, 2))
}
