package main

import (
	"fmt"
	"strconv"
)

func sumBase(n int, k int) int {
	temp, _ := strconv.ParseInt(strconv.FormatInt(int64(n), k), 0, 64)
	a, mySlice, sum := int(temp), make([]int, 0), 0
	for a != 0 {
		mySlice = append(mySlice, a%10)
		a /= 10
	}
	for _, v := range mySlice {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sumBase(34, 6))
}
