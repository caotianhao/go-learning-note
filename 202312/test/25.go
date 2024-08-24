package main

import (
	"fmt"
)

func percent(n, a, b int) float64 {
	p := float64(1) / float64(n+a+b)
	var s complex64 = complex(0, 1)
	var s1 complex64 = complex(0, -1)
	fmt.Println(s1 * s)
	return p
}

func main() {
	fmt.Println(percent(10, 17, 21)) // 0.62929
}
