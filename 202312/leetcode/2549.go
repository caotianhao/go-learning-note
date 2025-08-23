package main

import "fmt"

func distinctIntegers(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}

func main() {
	fmt.Println(distinctIntegers(1))
	fmt.Println(distinctIntegers(11))
	fmt.Println(distinctIntegers(29))
}
