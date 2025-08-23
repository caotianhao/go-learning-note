package main

import "fmt"

func generateTheString(n int) string {
	s := ""
	if n%2 == 1 {
		for i := 0; i < n; i++ {
			s += "a"
		}
	} else {
		for i := 0; i < n-1; i++ {
			s += "a"
		}
		s += "b"
	}
	return s
}

func main() {
	fmt.Println(generateTheString(18))
}
