package main

import "fmt"

func thunder3(n int) int {
	if n < 4 {
		return 1
	} else if n < 7 {
		return 3
	} else if n < 15 {
		return 7
	} else if n < 60 {
		return 205
	} else if n < 150 {
		return 717
	}
	return 5831
}

func main() {
	fmt.Println(thunder3(63))
}
