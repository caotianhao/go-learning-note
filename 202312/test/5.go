package main

import (
	"fmt"
)

const MOD = 1000000007

func countDeletionWays(nums []int, k int) int {
	n := len(nums)

	// 辅助函数：计算从前i个元素中删除k个元素的方式数
	var countWays func(int, int) int
	countWays = func(i, k int) int {
		if i < k {
			return 0
		}
		if k == 0 {
			return 1
		}

		// 选择删除第i个元素或不删除
		ways := countWays(i-1, k-1) + countWays(i-1, k)
		return ways % MOD
	}

	result := countWays(n, k)
	return result
}

func main() {
	nums := []int{1, 4, 2, 3, 6, 7}
	k := 1
	result := countDeletionWays(nums, k)
	fmt.Println(result) // 输出：0
}
