package main

func setZeros(matrix [][]int) {
	mySliceRow, mySliceCol := make([]int, 0), make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mySliceRow = append(mySliceRow, i)
				mySliceCol = append(mySliceCol, j)
			}
		}
	}
	for _, v := range mySliceRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v][i] = 0
		}
	}
	for _, v := range mySliceCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}
}

func main() {
	arr := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeros(arr)
}
