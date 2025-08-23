package main

import (
	"fmt"
)

func gardenNoAdj(n int, paths [][]int) []int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, path := range paths {
		adj[path[0]-1] = append(adj[path[0]-1], path[1]-1)
		adj[path[1]-1] = append(adj[path[1]-1], path[0]-1)
	}
	ans := make([]int, n)

	fmt.Println("ans", ans)
	for i := 0; i < n; i++ {
		colored := make([]bool, 5)
		for _, vertex := range adj[i] {
			colored[ans[vertex]] = true
		}
		fmt.Println(colored)
		for j := 1; j <= 4; j++ {
			if !colored[j] {
				ans[i] = j
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})) // 1 2 3
}
