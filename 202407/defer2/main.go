package main

import (
	"fmt"
	"os"
)

func defer2() {
	fmt.Print("111")
	defer fmt.Print("222")
	fmt.Print("333")
	os.Exit(1)
}

func main() {
	defer2() // 111333
}
