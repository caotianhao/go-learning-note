package main

import (
	"fmt"
	"math"
)

func findComplement(num int) int {
	for i := 0; i < 32; i++ {
		if num >= int(math.Pow(2.0, float64(i))) && num < int(math.Pow(2.0, float64(i+1))) {
			return num ^ (int(math.Pow(2.0, float64(i+1))) - 1)
		}
	}
	return 1
}

func main() {
	fmt.Println(findComplement(7))  //0
	fmt.Println(findComplement(5))  //2
	fmt.Println(findComplement(10)) //5
	fmt.Println(findComplement(1))  //0
	fmt.Println(findComplement(0))  //1
}
