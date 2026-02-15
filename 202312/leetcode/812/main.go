package main

import (
	"fmt"
	"math"
	"sort"
)

func largestTriangleArea(points [][]int) float64 {
	l, allS, a, b, c, s := len(points), make([]float64, 0), 0.0, 0.0, 0.0, 0.0
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				a = all2PointsFar(points[i], points[j])
				b = all2PointsFar(points[i], points[k])
				c = all2PointsFar(points[j], points[k])
				p := (a + b + c) / 2
				s = math.Sqrt(p * (p - a) * (p - b) * (p - c))
				allS = append(allS, s)
			}
		}
	}
	sort.Float64s(allS)
	return allS[len(allS)-1]
}

func all2PointsFar(point1 []int, point2 []int) float64 {
	a := math.Pow(float64(point1[0]-point2[0]), 2.0)
	b := math.Pow(float64(point1[1]-point2[1]), 2.0)
	return math.Sqrt(a + b)
}

func main() {
	fmt.Println(all2PointsFar([]int{0, 0}, []int{3, 4}))
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}))
}
