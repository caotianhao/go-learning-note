package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("exist.txt")
	fmt.Println(string(f), err)
}
