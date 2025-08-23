package main

import "fmt"

func longestString(x, y, z int) int {
	// memo 矩阵
	memo := make([][][][3]int, x+1)
	for i := range memo {
		memo[i] = make([][][3]int, y+1)
		for j := range memo[i] {
			memo[i][j] = make([][3]int, z+1)
			for k := range memo[i][j] {
				memo[i][j][k] = [3]int{-1, -1, -1}
			}
		}
	}
	var dfs func(int, int, int, int) int
	// flag : 0 AA    1 BB    2 AB
	dfs = func(x, y, z, flag int) (res int) {
		p := &memo[x][y][z][flag]
		if *p != -1 {
			return *p
		}
		// 如果前面结尾是 AA，那么后面只能是 BB
		if flag == 0 {
			if y > 0 {
				res = 2 + dfs(x, y-1, z, 1)
			}
		} else {
			// 如果前面结尾不是 AA，那么后面只能可能是 AA 或 AB
			// AB 情况
			if z > 0 {
				res = 2 + dfs(x, y, z-1, 2)
			}
			// AA 情况，并且和 AB 情况取最大值
			if x > 0 {
				res = max(res, 2+dfs(x-1, y, z, 0))
			}
		}
		*p = res // 记忆化搜索
		return   // 退出
	}
	// 上面已经有和 flag==2 的比较了，所以这里只比较 flag 为 0 和 1 的情况
	return max(dfs(x, y, z, 0), dfs(x, y, z, 1))
}

func main() {
	fmt.Println(longestString(2, 5, 1))
}
