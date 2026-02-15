package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	_, _ = fmt.Scanf("%f", &a)
	fmt.Printf("%.2f", math.Abs(a))
}
