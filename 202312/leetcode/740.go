package main

import "fmt"

func deleteAndEarn(nums []int) int {
	// 找出数组最大值
	maxE := -1
	for _, v := range nums {
		if v > maxE {
			maxE = v
		}
	}
	// 创建一个数组，下标的值代表下标在 nums 中的数量
	// 这样的话该题就变为了打家劫舍
	cnt, dp := make([]int, maxE+1), make([]int, maxE+2)
	for _, v := range nums {
		cnt[v]++
	}
	dp[1] = cnt[0]
	for i := 2; i < maxE+2; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+cnt[i-1]*(i-1))
	}
	return dp[maxE+1]
}

func main() {
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
	fmt.Println(deleteAndEarn([]int{2}))
}
