package main

import "fmt"

func checkString(s string) bool {
	flag := true
	for _, v := range s {
		if v == 'a' {
			if flag {
				continue
			} else {
				return false
			}
		}
		if v == 'b' {
			flag = false
		}
	}
	return true
}

func main() {
	fmt.Println(checkString("aaabbb"))
	fmt.Println(checkString("bbb"))
	fmt.Println(checkString("aaa"))
	fmt.Println(checkString("abaa"))
}
