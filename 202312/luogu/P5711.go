package main

import "fmt"

func main() {
	var year int
	_, _ = fmt.Scanf("%d", &year)
	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Printf("%d", 1)
		} else {
			if year%400 == 0 {
				fmt.Printf("%d", 1)
			} else {
				fmt.Printf("%d", 0)
			}
		}
	} else {
		fmt.Printf("%d", 0)
	}
}
