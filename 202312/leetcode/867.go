package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	l, l2, arrT := len(matrix), len(matrix[0]), make([][]int, 0)
	for j := 0; j < l2; j++ {
		colSlice := make([]int, 0)
		for i := 0; i < l; i++ {
			colSlice = append(colSlice, matrix[i][j])
		}
		arrT = append(arrT, colSlice)
	}
	return arrT
}

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	fmt.Println(transpose(arr))
}
