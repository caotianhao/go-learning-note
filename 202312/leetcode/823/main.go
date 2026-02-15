package main

import (
	"fmt"
	"sort"
)

func numFactoredBinaryTrees(nums []int) int {
	mod, answer := int64(1e9+7), int64(0)
	sort.Ints(nums)
	l := len(nums)
	dp := make([]int64, l)
	for i := 0; i < l; i++ {
		// 默认自己为根节点，且没有孩子算一种合法形式，所以初始化为 1
		dp[i] = 1
		// 双指针，因为排序完，所以乘积等于当前数的一定在其前面
		for left, right := 0, i-1; left <= right; left++ {
			// 如果大了就回去一个，直到 <=
			for left <= right && int64(nums[left])*int64(nums[right]) > int64(nums[i]) {
				right--
			}
			// 如果刚好相等，则看左右孩子是否相等
			if left <= right && int64(nums[left])*int64(nums[right]) == int64(nums[i]) {
				if left == right {
					// 若相等，根据乘法原理直接相乘
					dp[i] = (dp[i] + dp[left]*dp[right]) % mod
				} else {
					// 若不等，则还需乘 2
					dp[i] = (dp[i] + 2*dp[left]*dp[right]) % mod
				}
			}
		}
		answer = (answer + dp[i]) % mod
	}
	return int(answer)
}

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4}))           // 3
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))    // 7
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 8, 10})) // 12
}
