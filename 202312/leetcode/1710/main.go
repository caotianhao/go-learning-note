package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	l, size, sizeInd, thing, tmp := len(boxTypes), 0, 0, 0, 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for i := 0; i < l; i++ {
		size += boxTypes[i][0]
		if size > truckSize {
			sizeInd = i
			break
		}
	}
	if size < truckSize {
		for i := 0; i < l; i++ {
			thing += boxTypes[i][0] * boxTypes[i][1]
		}
		return thing
	}
	for i := 0; i < sizeInd; i++ {
		thing += boxTypes[i][0] * boxTypes[i][1]
		tmp += boxTypes[i][0]
	}
	thing += (truckSize - tmp) * boxTypes[sizeInd][1]
	return thing
}

func main() {
	//fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
	fmt.Println(maximumUnits([][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1},
		{2, 2}, {1, 3}, {2, 5}, {3, 2}}, 35))
}
