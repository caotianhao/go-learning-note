package main

import "fmt"

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 || n%3 == 0 || n%5 == 0 {
		if n%2 == 0 {
			n /= 2
		}
		if n%3 == 0 {
			n /= 3
		}
		if n%5 == 0 {
			n /= 5
		}
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isUgly(18))
}
