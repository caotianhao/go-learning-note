package main

import (
	"fmt"
)

func encrypt(msg string) (c string) {
	cipherList := map[byte]int{
		'a': 2, 'b': 2, 'c': 2,
		'd': 3, 'e': 3, 'f': 3,
		'g': 4, 'h': 4, 'i': 4,
		'j': 5, 'k': 5, 'l': 5,
		'm': 6, 'n': 6, 'o': 6,
		'p': 7, 'q': 7, 'r': 7, 's': 7,
		't': 8, 'u': 8, 'v': 8,
		'w': 9, 'x': 9, 'y': 9, 'z': 9,
	}
	for _, v := range msg {
		if v == 'A' {
			c += "9"
		} else if v >= 'B' && v <= 'Z' {
			c += fmt.Sprintf("%d", cipherList[byte(v+31)])
		} else if v >= 'a' && v <= 'z' {
			c += fmt.Sprintf("%d", cipherList[byte(v)])
		} else {
			c += string(v)
		}
	}
	return
}

func main() {
	var msg string
	_, _ = fmt.Scanf("%s", &msg)
	cipher := encrypt(msg)
	fmt.Println(cipher)
}
