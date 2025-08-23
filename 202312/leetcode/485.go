package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	cnt, maxN := 0, -1
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			if cnt > maxN {
				maxN = cnt
			}
			cnt = 0
		}
	}
	if cnt > maxN {
		maxN = cnt
	}
	return maxN
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
