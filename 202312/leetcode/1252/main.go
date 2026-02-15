package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
	targetArr := make([][]int, m)
	for i := 0; i < m; i++ {
		targetArr[i] = make([]int, n)
	}
	l := len(indices)
	cnt := 0
	for i := 0; i < l; i++ {
		for j := 0; j < n; j++ {
			targetArr[indices[i][0]][j]++
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < m; j++ {
			targetArr[j][indices[i][1]]++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if targetArr[i][j]%2 == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
	fmt.Println(oddCells(1, 1, [][]int{{0, 0}, {0, 0}}))
}
