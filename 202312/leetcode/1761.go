package main

import (
	"fmt"
	"math"
)

func minTrioDegree(n int, edges [][]int) int {
	matrix := make([][]int, n)
	degree := make([]int, n)
	minN := math.MaxInt64
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, v := range edges {
		matrix[v[0]-1][v[1]-1] = 1
		matrix[v[1]-1][v[0]-1] = 1
		degree[v[0]-1]++
		degree[v[1]-1]++
	}
	flag := true
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if matrix[i][j] == 1 && matrix[j][k] == 1 && matrix[k][i] == 1 {
					flag = false
					t := degree[i] + degree[j] + degree[k] - 6
					if t < minN {
						minN = t
					}
				}
			}
		}
	}
	if flag {
		return -1
	}
	return minN
}

func main() {
	fmt.Println(minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}))
	fmt.Println(minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}))
}
