package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	l, a := len(columnTitle), 0
	for i := 0; i < l; i++ {
		a += int(math.Pow(26.0, float64(i)) * float64(int(columnTitle[l-i-1])-64))
	}
	return a
}

func main() {
	fmt.Println(titleToNumber("CUETOV"))
}
