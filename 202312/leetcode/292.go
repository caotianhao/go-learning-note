package main

import "fmt"

func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canWinNim(4))
}
