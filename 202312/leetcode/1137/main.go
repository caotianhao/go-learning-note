package main

import "fmt"

func tribonacci(n int) int {
	answer := [38]int{}
	answer[0], answer[1], answer[2] = 0, 1, 1
	for i := 3; i < 38; i++ {
		answer[i] = answer[i-1] + answer[i-2] + answer[i-3]
	}
	return answer[n]
}

func main() {
	fmt.Println(tribonacci(22))
}
