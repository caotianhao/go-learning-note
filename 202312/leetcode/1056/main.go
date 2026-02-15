package main

import (
	"fmt"
	"math"
)

func confusingNumber(n int) bool {
	arr := make([]int, 0)
	n1 := n
	for n1 != 0 {
		if n1%10 == 2 || n1%10 == 3 || n1%10 == 4 || n1%10 == 5 || n1%10 == 7 {
			return false
		}
		arr = append(arr, n1%10)
		n1 /= 10
	}
	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] == 9 {
			arr[i] = 6
		} else if arr[i] == 6 {
			arr[i] = 9
		}
	}
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-1-i], arr[i]
	}
	newNum := 0
	for i := 0; i < l; i++ {
		newNum += arr[i] * int(math.Pow(10.0, float64(i)))
	}
	return n != newNum
}

func main() {
	fmt.Println(confusingNumber(6))
	fmt.Println(confusingNumber(89))
	fmt.Println(confusingNumber(11))
}
