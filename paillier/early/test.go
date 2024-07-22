package main

import "fmt"

func add(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Println(add(3, 45))
}
