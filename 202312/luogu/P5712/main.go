package main

import "fmt"

func main() {
	var apple int
	_, _ = fmt.Scanf("%d", &apple)
	if apple == 0 || apple == 1 {
		fmt.Printf("Today, I ate %d apple.", apple)
	} else {
		fmt.Printf("Today, I ate %d apples.", apple)
	}
}
