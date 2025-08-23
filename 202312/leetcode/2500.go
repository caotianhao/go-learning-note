package main

import (
	"fmt"
	"sort"
)

func deleteGreatestValue(grid [][]int) (sum int) {
	need := len(grid[0])
	for _, v := range grid {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}
	for need != 0 {
		for _, v := range matrixT(grid) {
			sum += lineMax(v)
			need--
		}
	}
	return
}

func lineMax(arr []int) (max int) {
	for _, v := range arr {
		if v >= max {
			max = v
		}
	}
	return
}

func matrixT(grid [][]int) [][]int {
	h, l := len(grid), len(grid[0])
	ret := make([][]int, l)
	for i := 0; i < l; i++ {
		ret[i] = make([]int, h)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			ret[j][i] = grid[i][j]
		}
	}
	return ret
}

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
	fmt.Println(deleteGreatestValue([][]int{{10}}))
	//fmt.Println(lineMax([]int{3, 3, 3, 1}))
	//fmt.Println(del2500([]int{3, 3, 3, 1}, 3))
}
