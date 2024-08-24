package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	_, _ = fmt.Scanf("%f %f %f", &a, &b, &c)
	delta := b*b - 4*a*c
	if delta < 0 {
		fmt.Printf("No answer!")
	} else if delta == 0 {
		fmt.Printf("x1=x2=%.5f", -b/2/a)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		if x1 > x2 {
			fmt.Printf("x1=%.5f;x2=%.5f", x2, x1)
		} else {
			fmt.Printf("x1=%.5f;x2=%.5f", x1, x2)
		}
	}
}
