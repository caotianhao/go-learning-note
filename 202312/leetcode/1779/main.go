package main

import (
	"fmt"
	"math"
)

func nearestValidPoint(x int, y int, points [][]int) int {
	validPoint, minDis, validPoint2 := make([][]int, 0), 20001, make([][]int, 0)
	for _, v := range points {
		if x == v[0] || y == v[1] {
			validPoint = append(validPoint, v)
		}
	}
	for _, v := range validPoint {
		t := int(math.Abs(float64(x-v[0]) + float64(y-v[1])))
		if t < minDis {
			minDis = t
		}
	}
	for _, v := range validPoint {
		if int(math.Abs(float64(x-v[0])+float64(y-v[1]))) == minDis {
			validPoint2 = append(validPoint2, v)
		}
	}
	if len(validPoint2) == 0 {
		return -1
	}
	for i := range points {
		if points[i][0] == validPoint2[0][0] && points[i][1] == validPoint2[0][1] {
			return i
		}
	}
	return 888
}

func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))
}
