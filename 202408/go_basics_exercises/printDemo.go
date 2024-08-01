package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Triple")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Quintuple")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("Triple-Quintuple")
		} else {
			fmt.Println(i)
		}
	}
}
