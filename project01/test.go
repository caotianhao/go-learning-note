package main

import (
	"fmt"
	"go-cth/project02"
	"time"
)

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d", &b)
	t0 := time.Now()
	project02.PrintPrime(a, b)
	t := time.Since(t0)
	fmt.Println(t)
}
