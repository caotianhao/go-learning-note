package main

import "fmt"

func generate(numRows int) [][]int {
	yh := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		yh[i] = make([]int, i+1)
		yh[i][0] = 1
		yh[i][i] = 1
		for j := 1; j < i; j++ {
			yh[i][j] = yh[i-1][j-1] + yh[i-1][j]
		}
	}
	return yh
}

func main() {
	//fmt.Println(generate(1))
	//fmt.Println(generate(2))
	//fmt.Println(generate(3))
	fmt.Println(generate(4))
}
