package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	num := n
	sum := 0
	bitNum := 0
	arr := make([]int, 0)
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
		bitNum++
	}
	for _, v := range arr {
		sum += myPow1134(v, bitNum)
	}
	return sum == num
}

func myPow1134(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(isArmstrong(561))
	fmt.Println(isArmstrong(153))
}
