package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Day(), t.Month(), t.Year())

	a := 0
	t2 := time.Now()
	for i := 0; i < 1e8; i++ {
		a++
	}
	t3 := time.Now()
	fmt.Println(t3.Sub(t2))
}
