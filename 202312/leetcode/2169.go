package main

import (
	"fmt"
)

//思想就是把辗转相除每一步的商相加
func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	} else if num1 == num2 {
		return 1
	}
	sum := 0
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	for ; num1%num2 != 0; num1, num2 = num2, num1%num2 {
		sum += num1 / num2
	}
	return sum + num1/num2
}

func main() {
	fmt.Println(countOperations(90, 46))
}
