package main

import "fmt"

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		if !help2133(matrix[i], n) {
			return false
		}
	}
	for i := 0; i < n; i++ {
		t := make([]int, 0)
		for j := 0; j < n; j++ {
			t = append(t, matrix[j][i])
		}
		if !help2133(t, n) {
			return false
		}
	}
	return true
}

func help2133(arr []int, n int) bool {
	m := map[int]struct{}{}
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkValid(nil))
}
