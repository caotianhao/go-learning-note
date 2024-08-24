package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p float64
	_, _ = fmt.Scanf("%f %f %f", &a, &b, &c)
	p = (a + b + c) / 2.0
	fmt.Printf("%.1f", math.Sqrt(p*(p-a)*(p-b)*(p-c)))
}
