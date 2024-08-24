package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d int
	_, _ = fmt.Scanf("%d %d\n%d %d", &a, &b, &c, &d)
	fmt.Printf("%.3f", math.Sqrt(math.Pow(float64(d)-float64(b), 2)+math.Pow(float64(c)-float64(a), 2)))
}
