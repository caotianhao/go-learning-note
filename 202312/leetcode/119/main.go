package main

import "fmt"

func getRow(rowIndex int) []int {
	yh := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		yh[i] = make([]int, i+1)
		yh[i][0], yh[i][i] = 1, 1
		for j := 1; j < i; j++ {
			yh[i][j] = yh[i-1][j-1] + yh[i-1][j]
		}
	}
	return yh[rowIndex]
}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
	fmt.Println(getRow(6))
}
