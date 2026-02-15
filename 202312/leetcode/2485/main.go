package main

import "fmt"

func pivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		l, r := 0, 0
		for j := 1; j <= i; j++ {
			l += j
		}
		for k := i; k <= n; k++ {
			r += k
		}
		if l == r {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
	for i := 1; i <= 3888; i++ {
		if pivotInteger(i) != -1 {
			fmt.Println(i)
		}
	}
}
