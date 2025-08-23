package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	l, l2, flag := len(matrix), len(matrix[0]), true
LOOP:
	for i := 0; i < l-1; i++ {
		for j := 0; j < l2-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				flag = false
				break LOOP
			}
		}
	}
	return flag
}

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(arr))
}
