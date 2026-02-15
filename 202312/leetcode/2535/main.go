package main

import (
	"fmt"
	"math"
)

func differenceOfSum(nums []int) int {
	sum, sum1 := 0, 0
	for _, v := range nums {
		sum += v
	}
	for _, v := range nums {
		sum1 += sumOfDivide(v)
	}
	return int(math.Abs(float64(sum - sum1)))
}

func sumOfDivide(n int) (sum int) {
	slice2535 := make([]int, 0)
	for n != 0 {
		slice2535 = append(slice2535, n%10)
		n /= 10
	}
	for _, v := range slice2535 {
		sum += v
	}
	return
}

func main() {
	fmt.Println(differenceOfSum([]int{1, 2, 3, 4, 5, 19}))
}
