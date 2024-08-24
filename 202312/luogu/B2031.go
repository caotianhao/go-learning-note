package main

import (
	"fmt"
	"math"
)

func getLength2031(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(y2-y1, 2) + math.Pow(x2-x1, 2))
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	_, _ = fmt.Scanf("%f %f %f %f %f %f", &x1, &y1, &x2, &y2, &x3, &y3)
	a := getLength2031(x1, y1, x2, y2)
	b := getLength2031(x2, y2, x3, y3)
	c := getLength2031(x1, y1, x3, y3)
	p := (a + b + c) / 2
	fmt.Printf("%.2f", math.Sqrt(p*(p-a)*(p-b)*(p-c)))
}
