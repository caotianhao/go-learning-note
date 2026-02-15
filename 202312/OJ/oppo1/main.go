package main

import (
	"fmt"
)

func getMaxContribute(nums []int) int {
	maxN := -1
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			con := (nums[i] + nums[j]) * min(j-i, l-j+i)
			if con > maxN {
				maxN = con
			}
		}
	}
	return maxN
}

func main() {
	n := 0
	a := 0
	arr := make([]int, 0)
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	fmt.Println(getMaxContribute(arr))
}
