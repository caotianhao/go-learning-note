package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	p, q, sum, ret := 0, 0, 0, 100001
	for q < len(nums) {
		sum += nums[q]
		for sum >= target {
			ret = min(q-p+1, ret)
			sum -= nums[p]
			p++
		}
		q++
	}
	if ret == 100001 {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
}
