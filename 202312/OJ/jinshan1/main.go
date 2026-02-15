package main

import (
	"fmt"
	"math"
)

func singSong8(nums []int) int {
	n, ans := len(nums), math.MinInt64
	left, right := 0, 1
	for right < n {
		if nums[right]-nums[right-1] > 8 {
			ans = max(ans, right-left)
			left = right
		}
		right++
	}
	ans = max(ans, right-left)
	return ans
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	m := 0
	music := make([]int, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &m)
		music = append(music, m)
	}
	fmt.Println(singSong8(music))
}
