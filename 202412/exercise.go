package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a, b := 3451, 15589
	c := gcd(a, b)
	fmt.Println(c)
	fmt.Printf("%d/%d\n", a/c, b/c)
}
