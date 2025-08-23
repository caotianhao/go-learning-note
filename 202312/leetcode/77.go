package main

import "fmt"

func combine(n int, k int) (res [][]int) {
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		if len(path)+(n-i+1) < k {
			return
		}
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		path = append(path, i)
		dfs(i + 1)
		path = path[:len(path)-1]
		dfs(i + 1)
	}
	dfs(1)
	return
}

func main() {
	fmt.Println(combine(15, 2))
}
