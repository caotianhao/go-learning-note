package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	l, s, flag := len(word), "", true
	for i := 0; i < l; i++ {
		if ch == word[i] {
			if i != l-1 {
				flag = false
				s += reverse2000(word[:i+1])
				s += word[i+1:]
				break
			} else {
				return reverse2000(word)
			}
		}
	}
	if flag {
		s = word
	}
	return s
}

func reverse2000(s string) string {
	str := []rune(s)
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str)
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	//fmt.Println(reverse2000("abcd"))
}
