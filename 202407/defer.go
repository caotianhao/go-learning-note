package main

import "fmt"

func dd() {
	i := 10
	defer fmt.Println("defer1", i)
	i++
	defer fmt.Println("defer2", i)
	fmt.Println(i)
}

func main() {
	dd() // 10
}
