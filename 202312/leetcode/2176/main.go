package main

import "fmt"

func countPairs(nums []int, k int) int {
	l, cnt := len(nums), 0
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			if j <= i {
				continue
			}
			if nums[i] == nums[j] && i*j%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	fmt.Println(countPairs(nums, k))
}
