package main

import (
	"fmt"
)

func numSpecial(mat [][]int) int {
	l, l2, cnt := len(mat), len(mat[0]), 0
	for i := 0; i < l; i++ {
		for j := 0; j < l2; j++ {
			if mat[i][j] == 1 && sumRow(mat, i) == 1 && sumCol(mat, j) == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func sumRow(arr [][]int, rowIndex int) int {
	l, sum := len(arr[0]), 0
	for i := 0; i < l; i++ {
		sum += arr[rowIndex][i]
	}
	return sum
}

func sumCol(arr [][]int, colIndex int) int {
	l, sum := len(arr), 0
	for i := 0; i < l; i++ {
		sum += arr[i][colIndex]
	}
	return sum
}

func main() {
	arr := [][]int{{0, 0, 0, 1}, {1, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numSpecial(arr))
}
