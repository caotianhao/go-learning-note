package main

import "fmt"

func defer1() {
	i := 10
	defer fmt.Printf("%d ", i)
	i++
	defer fmt.Printf("%d ", i)
	fmt.Printf("%d ", i)
}

func main() {
	defer1() // 11 11 10
}
