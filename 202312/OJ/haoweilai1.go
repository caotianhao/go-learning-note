package main

import "fmt"

func minOperation(s1, s2 string) int {
	n, m := len(s1), len(s2)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		v := &memo[i][j]
		if *v != -1 {
			return *v
		}
		if s1[i] == s2[j] {
			return dfs(i-1, j-1)
		}
		return 1 + min(min(dfs(i-1, j), dfs(i, j-1)), dfs(i-1, j-1))
	}
	return dfs(n-1, m-1)
}

func main() {
	var word1, word2 string
	_, _ = fmt.Scanf("%s", &word1)
	_, _ = fmt.Scanf("%s", &word2)
	fmt.Printf("%d", minOperation(word1, word2))
}
