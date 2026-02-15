package main

import "fmt"

type SubrectangleQueries struct {
	arr [][]int
}

func Constructor1476(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1, col1,
	row2, col2, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sq.arr[i][j] = newValue
		}
	}
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	return sq.arr[row][col]
}

func main() {
	sub := Constructor1476([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	fmt.Println(sub.GetValue(0, 2))
	sub.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Println(sub.GetValue(0, 2))
	fmt.Println(sub.GetValue(3, 1))
	sub.UpdateSubrectangle(3, 0, 3, 2, 10)
	fmt.Println(sub.GetValue(3, 1))
	fmt.Println(sub.GetValue(0, 2))
}
