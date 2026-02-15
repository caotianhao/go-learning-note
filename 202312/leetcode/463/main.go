package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	l, l2, theNumberOfOne := len(grid), len(grid[0]), 0
	for i := 0; i < l; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 {
				theNumberOfOne++
			}
		}
	}
	theNumberOfOne *= 4
	for i := 0; i < l-1; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 && grid[i+1][j] == 1 {
				theNumberOfOne--
			}
		}
	}
	for i := 1; i < l; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 && grid[i-1][j] == 1 {
				theNumberOfOne--
			}
		}
	}
	for j := 0; j < l2-1; j++ {
		for i := 0; i < l; i++ {
			if grid[i][j] == 1 && grid[i][j+1] == 1 {
				theNumberOfOne--
			}
		}
	}
	for j := 1; j < l2; j++ {
		for i := 0; i < l; i++ {
			if grid[i][j] == 1 && grid[i][j-1] == 1 {
				theNumberOfOne--
			}
		}
	}
	return theNumberOfOne
}

func main() {
	arr := [][]int{{0, 1, 0, 0, 1, 1, 1, 1, 0, 0}, {1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 1, 0, 0, 1, 0, 0, 1, 1, 1}, {1, 1, 0, 0, 0, 1, 0, 0, 1, 1}}
	fmt.Println(islandPerimeter(arr))
}
