package main

import "fmt"

func minOperations1551(n int) int {
	// 3 -> 2
	// 5 -> 4+2
	// 7 -> 6+4+2
	// ...
	// 2 -> 1
	// 4 -> 3+1
	// 6 -> 5+3+1
	return (n*n - (n & 1)) / 4
}

func main() {
	fmt.Println(minOperations1551(1))
}
