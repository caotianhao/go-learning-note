package main

import "fmt"

func closedIsland(grid [][]int) (cnt int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) bool
	dfs = func(x int, y int) bool {
		// 2. 如果碰壁，则显然不符合题意
		if x < 0 || x > m-1 || y < 0 || y > n-1 {
			return false
		}
		// 3. 如果不是 0，显然是被包围的
		if grid[x][y] != 0 {
			return true
		}
		// 4. 把本身位置标为 -1，其实只要不是 0 就行
		grid[x][y] = -1
		// 5. 向四周搜索
		ret1, ret2, ret3, ret4 := dfs(x-1, y), dfs(x+1, y), dfs(x, y-1), dfs(x, y+1)
		return ret1 && ret2 && ret3 && ret4
		// 这里不能直接这样 return，否则若第一个为 false 则直接为 false
		// 会少状态
		//return dfs(x-1, y) && dfs(x+1, y) && dfs(x, y-1) && dfs(x, y+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 1. 如果该位置为 0，则从其上下左右扩展搜索
			if grid[i][j] == 0 && dfs(i, j) {
				// 6. 如果被封闭，则 cnt++
				cnt++
			}
		}
	}
	return
}

func main() {
	fmt.Println(closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}))
	fmt.Println(closedIsland([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}))
}
