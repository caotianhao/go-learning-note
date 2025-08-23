package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen, secondLen int) (ans int) {
	// 分两种情况
	// 假设 firstLen 的数组在左/右面
	// 假设 firstLen 的数组在左面，那么从 firstLen 下标开始枚举 secondLen 的左端点
	// 使用一个变量维护左面的最大值，再计算右面的最大值，最后返回二者和的最大值
	// 要使用前缀和数组
	l := len(nums)
	// 前缀和数组
	prefix := make([]int, l+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	// leftMax 是左面最大值
	// 要熟悉前缀和数组的计算形式
	// 计算前缀和这种的，前后两个下标的差一定是区间的长度
	for i, leftMax := firstLen, 0; i+secondLen-1 < l; i++ {
		leftMax = max(leftMax, prefix[i]-prefix[i-firstLen])
		ans = max(ans, leftMax+prefix[i+secondLen]-prefix[i])
	}
	for i, rightMax := secondLen, 0; i+firstLen-1 < l; i++ {
		rightMax = max(rightMax, prefix[i]-prefix[i-secondLen])
		ans = max(ans, rightMax+prefix[i+firstLen]-prefix[i])
	}
	return
}

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3))
}
