package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	_, _ = fmt.Scanf("%f", &a)
	if a == 0 {
		fmt.Printf("%d", 0)
	} else if a < 0 {
		fmt.Printf("%d", int(math.Ceil(a)))
	} else {
		fmt.Printf("%d", int(math.Floor(a)))
	}
}
