package main

import (
	"fmt"
)

func findWords(words []string) []string {
	l, str := len(words), make([]string, 0)
	for i := 0; i < l; i++ {
		temp := line500(words[i][0])
		flag := true
		for j := 0; j < len(words[i]); j++ {
			if temp != line500(words[i][j]) {
				flag = false
			}
		}
		if flag == true {
			str = append(str, words[i])
		}
	}
	return str
}

func line500(c byte) int {
	if c == 'q' || c == 'w' || c == 'e' || c == 'r' || c == 't' ||
		c == 'y' || c == 'u' || c == 'i' || c == 'o' || c == 'p' ||
		c == 'Q' || c == 'W' || c == 'E' || c == 'R' || c == 'T' ||
		c == 'Y' || c == 'U' || c == 'I' || c == 'O' || c == 'P' {
		return 1
	} else if c == 'a' || c == 's' || c == 'd' || c == 'f' ||
		c == 'g' || c == 'h' || c == 'j' || c == 'k' || c == 'l' ||
		c == 'A' || c == 'S' || c == 'D' || c == 'F' ||
		c == 'G' || c == 'H' || c == 'J' || c == 'K' || c == 'L' {
		return 2
	} else {
		return 3
	}
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
