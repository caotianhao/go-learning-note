package main

import "fmt"

func commonFactors(a int, b int) int {
	minF, cnt := a, 0
	if b < a {
		minF = b
	}
	for i := 1; i <= minF; i++ {
		if a%i == 0 && b%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(commonFactors(12, 6))
}
