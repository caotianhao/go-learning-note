package main

import (
	"fmt"
	"strconv"
)

func findColumnWidth(grid [][]int) []int {
	row, col := len(grid), len(grid[0])
	res := make([]int, col)
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			res[i] = max(res[i], len(strconv.FormatInt(int64(grid[j][i]), 10)))
		}
	}
	return res
}

func main() {
	fmt.Println(findColumnWidth([][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}))
}
