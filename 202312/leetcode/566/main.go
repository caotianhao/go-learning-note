package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	l := len(mat)
	l2 := len(mat[0])
	temp := make([]int, 0)
	temp2 := make([]int, 0)
	newMat := make([][]int, 0)
	if l*l2 != r*c {
		return mat
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l2; j++ {
			temp = append(temp, mat[i][j])
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp2 = make([]int, 0)
			for k := 0; k < c; k++ {
				temp2 = append(temp2, temp[k+i*c])
			}
		}
		newMat = append(newMat, temp2)
	}
	return newMat
}

func main() {
	arr := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(arr, 1, 4))
}
