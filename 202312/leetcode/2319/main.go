package main

import "fmt"

func checkXMatrix(grid [][]int) bool {
	l, flag1, flag2 := len(grid), true, true
LOOP:
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j || i+j == l-1 {
				if grid[i][j] == 0 {
					flag1 = false
					break LOOP
				}
			} else {
				if grid[i][j] != 0 {
					flag2 = false
					break LOOP
				}
			}
		}
	}
	if flag1 && flag2 {
		return true
	} else {
		return false
	}
}

func main() {
	arr := [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}
	fmt.Println(checkXMatrix(arr))
}
