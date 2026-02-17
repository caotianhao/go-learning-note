package main

import "fmt"

// 此做法为 (n^3, 1) 做法
// 可使用哈希表优化为 (n^2, n^2)
func equalPairs(grid [][]int) (cnt int) {
	l := len(grid)
	ro := rotate2352(grid)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if isEqual2352(grid[i], ro[j]) {
				cnt++
			}
		}
	}
	return
}

func rotate2352(grid [][]int) [][]int {
	l := len(grid)
	res := make([][]int, l)
	for i := 0; i < l; i++ {
		res[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			res[j][i] = grid[i][l-1-j]
		}
	}
	return res
}

func isEqual2352(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
}
