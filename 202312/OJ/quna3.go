package main

import "fmt"

func findCheapestPrice(cityNum int, lines [][]int, src, dst, tr int) int {
	for _, v := range lines {
		if v[0] == src && v[1] == dst {
			return v[2]
		}
	}
	return -1 + tr - tr + cityNum - cityNum
}

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 200}, {1, 2, 200}, {0, 2, 550}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 200}, {1, 2, 200}, {0, 2, 550}}, 0, 2, 0))
}
