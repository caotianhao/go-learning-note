package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scanf("%d", &a)
	if a%3 == 0 && a%5 == 0 && a%7 == 0 {
		fmt.Printf("3 5 7")
	} else if a%3 == 0 && a%5 == 0 && a%7 != 0 {
		fmt.Printf("3 5")
	} else if a%3 == 0 && a%5 != 0 && a%7 == 0 {
		fmt.Printf("3 7")
	} else if a%3 != 0 && a%5 == 0 && a%7 == 0 {
		fmt.Printf("5 7")
	} else if a%3 == 0 && a%5 != 0 && a%7 != 0 {
		fmt.Printf("3")
	} else if a%3 != 0 && a%5 == 0 && a%7 != 0 {
		fmt.Printf("5")
	} else if a%3 != 0 && a%5 != 0 && a%7 == 0 {
		fmt.Printf("7")
	} else {
		fmt.Printf("n")
	}
}
