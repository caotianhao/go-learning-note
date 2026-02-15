package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	fmt.Printf("%d\n", rand.Intn(40))
}
