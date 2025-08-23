package main

import (
	"fmt"
)

func permutation(s string) (ans []string) {
	l := len(s)
	path := ""
	used := make([]bool, l)
	//var dfs func(int, string)
	var dfs func(int)
	//dfs = func(i int, path string) {
	dfs = func(i int) {
		if i == l {
			ans = append(ans, path)
			return
		}
		for j, is := range used {
			if !is {
				path += string(s[j])
				used[j] = true
				dfs(i + 1)
				used[j] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Println(permutation("pum"))
}
