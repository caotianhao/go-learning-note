package main

import "fmt"

func queensAttacktheKing(queens [][]int, king []int) (r [][]int) {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	quMap := make(map[[2]int]struct{})
	for _, q := range queens {
		quMap[[2]int{q[0], q[1]}] = struct{}{}
	}
	for i := 0; i < 8; i++ {
		p, q := king[0], king[1]
		for isIn1222(p, q) {
			if _, ok := quMap[[2]int{p + dir[i][0], q + dir[i][1]}]; ok {
				r = append(r, []int{p + dir[i][0], q + dir[i][1]})
				break
			}
			p, q = p+dir[i][0], q+dir[i][1]
		}
	}
	return
}

func isIn1222(x, y int) bool {
	return x >= 0 && x <= 7 && y >= 0 && y <= 7
}

func main() {
	fmt.Println(queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0}))
}
