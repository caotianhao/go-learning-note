package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(cur int, sum int) {
		if cur >= len(candidates) {
			return
		}
		if sum == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(cur+1, sum)
		if sum-candidates[cur] >= 0 {
			path = append(path, candidates[cur])
			dfs(cur, sum-candidates[cur])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return
}

func main() {
	fmt.Println(combinationSum([]int{2, 1, 6, 7}, 7))
}
