package main

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	l := len(s)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		row, col := startPos[0], startPos[1]
		do := s[i:]
		for j := 0; j < l-i; j++ {
			if do[j] == 'L' {
				col--
				if isInBound(row, col, n) {
					res[i]++
				} else {
					break
				}
			} else if do[j] == 'R' {
				col++
				if isInBound(row, col, n) {
					res[i]++
				} else {
					break
				}
			} else if do[j] == 'U' {
				row--
				if isInBound(row, col, n) {
					res[i]++
				} else {
					break
				}
			} else {
				row++
				if isInBound(row, col, n) {
					res[i]++
				} else {
					break
				}
			}
		}
	}
	return res
}

func isInBound(left, right, n int) bool {
	return left >= 0 && left < n && right < n && right >= 0
}

func main() {
	fmt.Println(executeInstructions(3, []int{0, 1}, "RRDDLU")) // 1,5,4,3,1,0
}
