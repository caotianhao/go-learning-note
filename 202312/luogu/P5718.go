package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a int
	minNow := math.MaxInt64
	_, _ = fmt.Scan(&n)
	if n == 0 {
		fmt.Printf("%d", 0)
		return
	}
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a)
		if a < minNow {
			minNow = a
		}
	}
	fmt.Printf("%d", minNow)
}
