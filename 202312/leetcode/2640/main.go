package main

import "fmt"

func findPrefixScore(nums []int) []int64 {
	l := len(nums)
	maxN, sum := make([]int, l), make([]int, l)
	maxN[0] = nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > maxN[i-1] {
			maxN[i] = nums[i]
		} else {
			maxN[i] = maxN[i-1]
		}
	}
	for i := 0; i < l; i++ {
		sum[i] = maxN[i] + nums[i]
	}
	res := make([]int64, l+1)
	for i, v := range sum {
		res[i+1] = res[i] + int64(v)
	}
	return res[1:]
}

func main() {
	fmt.Println(findPrefixScore([]int{2, 3, 7, 5, 10}))
}
