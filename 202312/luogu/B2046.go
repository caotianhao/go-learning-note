package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a == 100 {
		fmt.Printf("All")
	} else if a < 100 && a >= 0 {
		fmt.Printf("Walk")
	} else {
		fmt.Printf("Bike")
	}
}
