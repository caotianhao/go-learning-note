package main

import "fmt"

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println(restoreMatrix([]int{3, 8}, []int{4, 7}))
	fmt.Println(restoreMatrix([]int{0}, []int{0}))
}
