package main

import "fmt"

func projectionArea(grid [][]int) int {
	l, xy, yz, xz, col := len(grid), 0, 0, 0, make([]int, 0)
	//xy：z 俯视，不为 0 的块数
	for _, v := range grid {
		for _, v2 := range v {
			if v2 != 0 {
				xy++
			}
		}
	}
	//xz：y 俯视，矩阵每一行的最大值之和
	for i := 0; i < l; i++ {
		xz += getMax883(grid[i])
	}
	//yz：x 俯视，矩阵每一列的最大值之和
	for i := 0; i < l; i++ {
		col = make([]int, 0)
		for j := 0; j < l; j++ {
			col = append(col, grid[j][i])
		}
		yz += getMax883(col)
	}
	return xy + xz + yz
}

func getMax883(arr []int) int {
	maxR := -1
	for _, v := range arr {
		if v > maxR {
			maxR = v
		}
	}
	return maxR
}

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}))                  //17
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}}))                  //8
	fmt.Println(projectionArea([][]int{{2}}))                             //5
	fmt.Println(projectionArea([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) //51
}
