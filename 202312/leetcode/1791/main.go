package main

import "fmt"

func findCenter(edges [][]int) int {
	if edges[1][0] == edges[0][0] || edges[1][0] == edges[0][1] {
		return edges[1][0]
	} else {
		return edges[1][1]
	}
}

func findCenter1(edges [][]int) int {
	a, b := edges[0][0], edges[0][1]
	c, d := edges[1][0], edges[1][1]
	if a == c || a == d {
		return a
	}
	return b
}

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {3, 1}, {1, 4}}))
	fmt.Println(findCenter1([][]int{{1, 2}, {3, 1}, {1, 4}}))
}
