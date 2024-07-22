package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	fmt.Println(rand.Prime(rand.Reader, 66))
}
