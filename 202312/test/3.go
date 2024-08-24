package main

import "fmt"

func backtracking(n, sum, start int, res []int) []int {
	if sum == 0 {
		// 找到一个满足条件的组合
		return res
	}
	for i := start; i <= n; i++ {
		if i <= sum {
			// 将当前数字添加到组合中
			res = append(res, i)
			// 递归查找下一个数字
			backtracking(n, sum-i, i+1, res)
			// 移除最后一个数字，继续搜索其他组合
			res = res[:len(res)-1]
		}
	}
	return res
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	// 计算从 1-n 的全排列，要满足题意的话，显然从 1 加到 n 的和必须为偶数
	sum := n * (n + 1) / 2
	if sum&1 == 1 {
		// 如果 sum 奇数，返回 -1
		fmt.Printf("-1")
	} else {
		// 然后选一些数，使它们的和为 sum/2，即可
		fmt.Println(backtracking(n, sum/2, 1, []int{}))
	}
}
