package main

import (
	"fmt"
	"math"
)

func bitwiseComplement(n int) int {
	for i := 0; i < 32; i++ {
		if n >= int(math.Pow(2.0, float64(i))) && n < int(math.Pow(2.0, float64(i+1))) {
			return n ^ (int(math.Pow(2.0, float64(i+1))) - 1)
		}
	}
	return 1
}

func main() {
	fmt.Println(bitwiseComplement(7))
	//fmt.Println(~5)
}
