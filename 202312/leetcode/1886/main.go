package main

import "fmt"

func findRotation(mat, target [][]int) bool {
	b := isEqual1886(mat, target)

	mat = rotate90(mat)
	b90 := isEqual1886(mat, target)

	mat = rotate90(mat)
	b180 := isEqual1886(mat, target)

	mat = rotate90(mat)
	b270 := isEqual1886(mat, target)

	return b || b90 || b180 || b270
}

func rotate90(matrix [][]int) [][]int {
	l := len(matrix)
	res := make([][]int, l)
	for i := 0; i < l; i++ {
		res[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			res[j][i] = matrix[l-1-i][j]
		}
	}
	return res
}

func isEqual1886(a, b [][]int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(findRotation([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}})) //t
	fmt.Println(findRotation([][]int{{1, 1}, {0, 1}}, [][]int{{1, 1}, {1, 0}}))                                   //t
	fmt.Println(findRotation([][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}))                                   //f
	//arr := [][]int{{1, 1}, {0, 1}}
	//arr1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(findRotation(arr1, [][]int{{1, 1}, {1, 0}}))
	//fmt.Println(reverse1886([]int{1, 2, 3, 6, 9, 11}))
	//fmt.Println("90", rotate90([][]int{{1, 1}, {0, 1}}))
	//fmt.Println("180", rotate180([][]int{{1, 1}, {0, 1}}))
	//fmt.Println("270", rotate270([][]int{{1, 1}, {0, 1}}))
	//fmt.Println("90", rotate90([][]int{{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}))
}
