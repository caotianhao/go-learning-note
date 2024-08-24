package main

import (
	"fmt"
	"sort"
)

func gcd1888(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd1888(b, a%b)
}

func main() {
	a, b, c := 0, 0, 0
	h := make([]int, 0)
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	h = append(h, a, b, c)
	sort.Ints(h)
	g := gcd1888(h[0], h[2])
	fmt.Printf("%d/%d", h[0]/g, h[2]/g)
}
