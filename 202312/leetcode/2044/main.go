package main

import "fmt"

func countMaxOrSubsets(nums []int) (cnt int) {
	l := len(nums)
	maxN := -1
	for i := 1; i < 1<<l; i++ {
		res := 0
		for j := 0; j < l; j++ {
			if i>>j&1 == 1 {
				res |= nums[l-j-1]
			}
		}
		if res > maxN {
			maxN = res
		}
	}
	for i := 1; i < 1<<l; i++ {
		res := 0
		for j := 0; j < l; j++ {
			if i>>j&1 == 1 {
				res |= nums[l-j-1]
			}
		}
		if res == maxN {
			cnt++
		}
	}
	return
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5})) // 6
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))    // 7
}
