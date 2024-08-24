package main

import (
	"fmt"
)

func makePlant(x, y, z int) (day int) {
	for i := 0; z > 0; i++ {
		z -= x
		if i%3 == 0 {
			z -= y
		}
		day++
	}
	return
}

func main() {
	x, y, z := 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d", &x, &y, &z)
	fmt.Printf("%d", makePlant(x, y, z))
}
