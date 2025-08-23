package main

import "fmt"

func permute(nums []int) (ans [][]int) {
	// 回溯入门
	// 先取到数组长度
	l := len(nums)
	// used 为标记该位置是否使用过
	// arr 为临时数组或者说是每一次的答案
	used, arr := make([]bool, l), make([]int, l)
	var dfs func(int)
	dfs = func(k int) {
		// 从 0 开始如果 k 达到了长度
		// 那就把数组添加到答案里
		if k == l {
			// 不能用 ans = append(ans, arr)
			// 要使用它的拷贝
			ans = append(ans, append([]int(nil), arr...))
			return
		}
		// 遍历 used 数组
		for ind, b := range used {
			// 如果为 false 就说明该位置的数字没有使用过
			if !b {
				// 就将其加入 arr
				arr[k] = nums[ind]
				// 置为 true
				used[ind] = true
				// 递归 k+1
				dfs(k + 1)
				// 使用完将其置为 false
				used[ind] = false
			}
		}
	}
	// 从 0 开始，就像树从根节点开始一样
	dfs(0)
	return
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
