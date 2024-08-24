package main

import "fmt"

func main() {
	var l, cnt int
	_, _ = fmt.Scanf("%d", &l)
	for l != 1 {
		l /= 2
		cnt++
	}
	fmt.Printf("%d", cnt+1)
}
