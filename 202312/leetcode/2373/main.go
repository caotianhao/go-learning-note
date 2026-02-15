package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	l := len(grid)
	ret := make([][]int, l-2)
	for i := 0; i < l-2; i++ {
		ret[i] = make([]int, l-2)
	}
	for i := 1; i < l-1; i++ {
		for j := 1; j < l-1; j++ {
			ret[i-1][j-1] = partMax2373(grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1],
				grid[i][j-1], grid[i][j], grid[i][j+1],
				grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1])
		}
	}
	return ret
}

func partMax2373(a ...int) int {
	maxN := 0
	for _, v := range a {
		if v > maxN {
			maxN = v
		}
	}
	return maxN
}

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}
