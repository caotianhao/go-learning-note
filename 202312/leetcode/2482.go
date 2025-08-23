package main

import "fmt"

func onesMinusZeros(grid [][]int) [][]int {
	row, col := len(grid), len(grid[0])
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
	}
	onesRow, zerosRow := make([]int, row), make([]int, row)
	onesCol, zerosCol := make([]int, col), make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				onesRow[i]++
				onesCol[j]++
			} else {
				zerosRow[i]++
				zerosCol[j]++
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[i][j] = onesRow[i] + onesCol[j] - zerosCol[j] - zerosRow[i]
		}
	}
	return res
}

func main() {
	fmt.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
}
