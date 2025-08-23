package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	tmp := divide1455(sentence)
	l := len(tmp)
	for i := 0; i < l; i++ {
		if strings.HasPrefix(tmp[i], searchWord) {
			return i + 1
		}
	}
	return -1
}

func divide1455(str string) (slice1455 []string) {
	s := ""
	for _, v := range str {
		if v == ' ' {
			slice1455 = append(slice1455, s)
			s = ""
			continue
		}
		s += string(v)
	}
	slice1455 = append(slice1455, s)
	return
}

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
	fmt.Println(divide1455("i love eating burger"))
}
