package main

import "fmt"

func combine080(n int, k int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(cur int) {
		if len(path)+(n-cur+1) < k {
			return
		}
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		path = append(path, cur)
		dfs(cur + 1)
		path = path[:len(path)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}

func main() {
	fmt.Println(combine080(5, 4))
}
