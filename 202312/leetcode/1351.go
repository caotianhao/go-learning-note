package main

import "fmt"

func countNegatives(grid [][]int) int {
	cnt, l, l2 := 0, len(grid), len(grid[0])
	for i := 0; i < l; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] < 0 {
				cnt += l2 - j
				break
			}
		}
	}
	return cnt
}

func main() {
	arr := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(arr))
}
