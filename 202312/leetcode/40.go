package main

import "fmt"

func combinationSum2(candidates []int, target int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(idx int, sum int) {
		if idx == len(candidates) {
			return
		}
		if sum == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		if sum-candidates[idx] >= 0 {
			path = append(path, candidates[idx])
			dfs(idx+1, sum-candidates[idx])
			path = path[:len(path)-1]
		}
		dfs(idx+1, sum)
	}
	dfs(0, target)
	return
}

func main() {
	fmt.Println(combinationSum2([]int{2, 1, 6, 7}, 7))
}
