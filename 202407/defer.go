package main

import "fmt"

func dd() {
	i := 10
	defer fmt.Println(i)
	i++
}

func main() {
	dd() // 10
}
