package main

import "fmt"

//m 行 n 列
//位于 grid[i][j] 的元素将会移动到 grid[i][j+1]
//位于 grid[i][n-1] 的元素将会移动到 grid[i+1][0]
//位于 grid[m-1][n-1] 的元素将会移动到 grid[0][0]
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m)
	if k == 0 {
		return grid
	}
	for i := 0; i < k; i++ {
		ret = make([][]int, m)
		for j := 0; j < m; j++ {
			ret[j] = make([]int, n)
		}
		ret[0][0] = grid[m-1][n-1]
		for a := 0; a < m-1; a++ {
			ret[a+1][0] = grid[a][n-1]
		}
		for a := 0; a < m; a++ {
			for b := 0; b < n-1; b++ {
				ret[a][b+1] = grid[a][b]
			}
		}
		grid = ret
	}
	return ret
}

func main() {
	fmt.Println(shiftGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, 1))
	fmt.Println(shiftGrid([][]int{
		{3, 8, 1, 9},
		{19, 7, 2, 5},
		{4, 6, 11, 10},
		{12, 0, 21, 13},
	}, 4))
	fmt.Println(shiftGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, 9))
	fmt.Println(shiftGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, 0))
}
