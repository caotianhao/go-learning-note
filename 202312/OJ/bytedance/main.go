package main

import (
	"fmt"
	"strconv"
)

func findNextInteger(digits []string, n int) int {
	return backtrack(digits, n, "")
}

func backtrack(digits []string, n int, current string) int {
	num, _ := strconv.Atoi(current)
	if num >= n {
		return num
	}
	maxNum := -1
	for _, digit := range digits {
		newNum := backtrack(digits, n, current+digit)
		if newNum >= n && (maxNum == -1 || newNum < maxNum) {
			maxNum = newNum
		}
	}
	return maxNum
}

func main() {
	digits := []string{"1", "3", "5"}
	n := 1630
	nextInteger := findNextInteger(digits, n)
	fmt.Println(nextInteger)
}
