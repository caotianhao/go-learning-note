package main

import "fmt"

func brokenMicroServer(nodes, number int, e [][]int) (cnt int) {
	edges := make([][]int, 0)
	for i := 0; i < nodes; i++ {
		if len(e[i]) == 1 && e[i][0] == -1 {
			continue
		}
		for _, quan := range e[i] {
			edges = append(edges, []int{i, quan - 1})
		}
	}
	for dst := 0; dst < nodes; dst++ {
		if validPath(nodes, edges, number-1, dst) {
			cnt++
		}
	}
	return
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	for _, e := range edges {
		p[find(e[0])] = find(e[1])
	}
	return find(source) == find(destination)
}

func main() {
	fmt.Println(brokenMicroServer(6, 3, [][]int{{2, 3}, {3}, {-1}, {5}, {1, 2}, {-1}}))
	fmt.Println(brokenMicroServer(3, 3, [][]int{{2}, {-1}, {-1}}))
}
