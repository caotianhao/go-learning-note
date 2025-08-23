package main

import "fmt"

func differenceOfSums(n int, m int) int {
	can, cant := 0, 0
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			can += i
		} else {
			cant += i
		}
	}
	return cant - can
}

func main() {
	fmt.Println(differenceOfSums(10, 3)) // 19
}
