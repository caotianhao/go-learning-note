package main

import "fmt"

func calculate(s string) int {
	str, x, y := []rune(s), 1, 0
	for i := 0; i < len(str); i++ {
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
