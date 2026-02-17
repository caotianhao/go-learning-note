package main

import "fmt"

func calculate(s string) int {
	str, x, y := []rune(s), 1, 0
	for i := range str {
		if str[i] == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}

func main() {
	fmt.Println(calculate("ABABABBBABABABABBBABAAAA"))
}
