package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	l := len(grid)
	retMatrix := make([][]int, l)
	for i := 0; i < l; i++ {
		retMatrix[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			rowMax, columnMax := -1, -1
			for k := 0; k < l; k++ {
				if rowMax <= grid[i][k] {
					rowMax = grid[i][k]
				}
			}
			for a := 0; a < l; a++ {
				if columnMax <= grid[a][j] {
					columnMax = grid[a][j]
				}
			}
			if grid[i][j] == rowMax {
				retMatrix[i][j] = rowMax
			}
			if grid[i][j] == columnMax {
				retMatrix[i][j] = columnMax
			}
			if grid[i][j] != rowMax && grid[i][j] != columnMax {
				retMatrix[i][j] = min(columnMax, rowMax)
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			ans += retMatrix[i][j] - grid[i][j]
		}
	}
	return
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}))
	fmt.Println(maxIncreaseKeepingSkyline([][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}))
}
