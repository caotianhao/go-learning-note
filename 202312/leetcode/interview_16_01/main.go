package main

import "fmt"

func swapNumbers(numbers []int) []int {
	a, b := numbers[0], numbers[1]
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return []int{a, b}
}

func main() {
	fmt.Println(swapNumbers([]int{2, 3}))
}
