package main

import "fmt"

func isPalindrome9(x int) bool {
	flag := true
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		str := fmt.Sprintf("%d", x)
		l := len(str)
		for i := 0; i < l; i++ {
			if str[i] != str[l-1-i] {
				flag = false
				break
			}
		}
	}
	return flag
}

func main() {
	fmt.Println(isPalindrome9(2255222))
}
