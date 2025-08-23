package main

import "fmt"

func rob(nums []int) int {
	l := len(nums)
	// 创建一个长度为 size+1 的数组，dp[0] 没啥用
	dp := make([]int, l+1)
	// 如果只有一间房，那么答案就是这间房里的钱
	dp[1] = nums[0]
	for i := 2; i < l+1; i++ {
		// 如果打劫最后一间房，那就不能打劫倒数第二间
		// 那么就是最后一间房的钱加上之前除了倒数第二间的值
		// 如果不打劫最后一间房，那么就是倒数第二间的值
		// 显然该取最大值
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	// dp 的长度为 l+1，所以返回 dp[l]
	return dp[l]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
