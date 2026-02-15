package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	isNeg := false
	if x < 0 {
		x = -x
		isNeg = true
	}
	digit := make([]int, 0)
	for x != 0 {
		digit = append(digit, x%10)
		x /= 10
	}
	l, res := len(digit), 0
	for i := 0; i < l; i++ {
		res += digit[i] * int(math.Pow(10.0, float64(l-1-i)))
	}
	if res <= -int(math.Pow(2, 31)) || res >= int(math.Pow(2, 31))-1 {
		return 0
	} else if isNeg {
		return -res
	}
	return res
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}
