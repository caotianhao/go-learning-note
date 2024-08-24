package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand25() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 25 {
			return 1 + (idx-1)%25
		}
	}
}

func rand7() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func main() {
	fmt.Println(rand25())
}
