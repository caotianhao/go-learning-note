package main

import "fmt"

func findTargetIn2DPlants(plants [][]int, target int) bool {
	m := len(plants)
	if m == 0 {
		return false
	}
	n := len(plants[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if plants[x][y] == target {
			return true
		} else if plants[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func main() {
	fmt.Println(findTargetIn2DPlants([][]int{{1}}, 2))
	fmt.Println(findTargetIn2DPlants([][]int{{1}}, 1))
	fmt.Println(findTargetIn2DPlants(nil, 1))
}
