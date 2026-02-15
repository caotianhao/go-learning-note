package main

import (
	"fmt"
	"math"
)

func floyd(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for start := 0; start < n; start++ {
				matrix[i][j] = minFloyd(matrix[i][j], matrix[i][start]+matrix[start][j])
			}
		}
	}
	return matrix
}

func minFloyd(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{0, 6, 13},
		{10, 0, 4},
		{5, math.MaxInt32, 0},
	}
	fmt.Println(floyd(matrix))
}
