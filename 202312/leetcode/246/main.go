package main

import "fmt"

func isStrobogrammatic(num string) bool {
	l := len(num)
	n := []byte(num)
	for i := 0; i < l/2; i++ {
		n[i], n[l-1-i] = n[l-i-1], n[i]
	}
	for i := 0; i < l; i++ {
		if num[i] == '2' || num[i] == '3' || num[i] == '4' || num[i] == '5' || num[i] == '7' {
			return false
		} else if num[i] == '0' && n[i] != '0' {
			return false
		} else if num[i] == '1' && n[i] != '1' {
			return false
		} else if num[i] == '6' && n[i] != '9' {
			return false
		} else if num[i] == '8' && n[i] != '8' {
			return false
		} else if num[i] == '9' && n[i] != '6' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isStrobogrammatic("5456132"))
	fmt.Println(isStrobogrammatic("69869"))
}
