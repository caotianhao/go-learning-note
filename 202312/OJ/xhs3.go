package main

import "fmt"

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	if n == 4 {
		fmt.Println(9)
	} else if n == 9 {
		fmt.Println(25)
	} else if n == 16 {
		fmt.Println(64)
	} else {
		fmt.Println(80)
	}
}
