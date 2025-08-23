package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)
	if l == 2 {
		return true
	}
	var k, b float64
	if coordinates[1][0] == coordinates[0][0] {
		for i := 2; i < l; i++ {
			if coordinates[i][0] != coordinates[0][0] {
				return false
			}
		}
		return true
	} else {
		k = float64(coordinates[1][1]-coordinates[0][1]) /
			float64(coordinates[1][0]-coordinates[0][0])
		b = float64(coordinates[0][1]) - k*float64(coordinates[0][0])
		for i := 2; i < l; i++ {
			if k*float64(coordinates[i][0])+b != float64(coordinates[i][1]) {
				return false
			}
		}
		return true
	}
}

func main() {
	//fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	//fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	fmt.Println(checkStraightLine([][]int{{0, 0}, {0, 1}, {0, -1}}))
}
