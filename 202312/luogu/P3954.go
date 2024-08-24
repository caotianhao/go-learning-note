package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Printf("%d", a*2/10+b*3/10+c*5/10)
}
