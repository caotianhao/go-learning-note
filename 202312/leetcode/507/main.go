package main

import (
	"fmt"
	"math"
)

// 如果它和除了它自身以外的所有正因子之和相等
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i
			sum += num / i
		}
	}
	return sum == num
}

func main() {
	//fmt.Println(checkPerfectNumber(199))
	fmt.Println(checkPerfectNumber(28))
	//fmt.Println(checkPerfectNumber(7))
	//fmt.Println(checkPerfectNumber(1))
}
