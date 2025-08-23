package main

import "fmt"

func diagonalSum(mat [][]int) int {
	sum, l := 0, len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j || i+j == l-1 {
				sum += mat[i][j]
			}
		}
	}
	return sum
}

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(diagonalSum(arr))
}
