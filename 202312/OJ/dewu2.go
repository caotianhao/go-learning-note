package main

import (
	"fmt"
	"math"
	"sort"
)

func theLeastNum(nums []int, target int) int {
	sort.Ints(nums)
	t := 0
	ans := math.MaxInt64
	if nums[len(nums)-1] > target {
		return ans
	}
	var dfs func(int, int)
	dfs = func(idx int, target int) {
		if target == 0 {
			if t < ans {
				ans = t
			}
		}
		for i := idx; i < len(nums); i++ {
			if nums[i] > target {
				return
			}
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			if t >= ans {
				return
			}
			t++
			dfs(i+1, target-nums[i])
			t--
		}
	}
	dfs(0, target)
	return ans
}

func main() {
	n, m := 0, 0
	a := 0
	arr := make([]int, 0)
	_, _ = fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &a)
		arr = append(arr, a)
	}
	ans := theLeastNum(arr, m)
	if ans != math.MaxInt64 {
		fmt.Printf("%d", ans)
	} else {
		fmt.Printf("No solution")
	}
}
