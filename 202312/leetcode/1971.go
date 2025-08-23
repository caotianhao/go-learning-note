package main

import "fmt"

// dfs
func validPath1(n int, edges [][]int, source, destination int) bool {
	adj := make([][]int, n)
	vis := make([]bool, n)
	for _, v := range edges {
		a, b := v[0], v[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	var dfs func(int, [][]int, []bool) bool
	dfs = func(s int, adj [][]int, vis []bool) bool {
		if s == destination {
			return true
		}
		vis[s] = true
		for _, v := range adj[s] {
			if !vis[v] && dfs(v, adj, vis) {
				return true
			}
		}
		return false
	}
	return dfs(source, adj, vis)
}

// 并查集
func validPath2(n int, edges [][]int, source, destination int) bool {
	// 1. init
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	// 2. find
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	// 3. edge
	for _, e := range edges {
		p[find(e[0])] = find(e[1])
	}
	// 4. judge
	return find(source) == find(destination)
}

func main() {
	fmt.Println(validPath1(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath2(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}
