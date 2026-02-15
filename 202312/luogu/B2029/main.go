package main

import (
	"fmt"
	"math"
)

func main() {
	var h, r int
	_, _ = fmt.Scanf("%d %d", &h, &r)
	v := 3.14 * float64(h) * float64(r) * float64(r) / 1000.0
	fmt.Printf("%d", int(math.Ceil(20.0/v)))
}
