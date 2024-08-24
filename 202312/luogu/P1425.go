package main

import "fmt"

func main() {
	var a, b, c, d int
	_, _ = fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	t := (c*60 + d) - (a*60 + b)
	fmt.Printf("%d %d", t/60, t%60)
}
