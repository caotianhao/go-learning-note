package main

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {
	pwdMap, encode, trueEncode := make(map[string]byte), make([]string, 0), ""
	var a byte = 'a'
	for i := 0; i < len(key); i++ {
		if string(key[i]) != " " {
			pwdMap[string(key[i])] = a
			a++
			key = strings.Replace(key, string(key[i]), " ", -1)
		}
	}
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			encode = append(encode, " ")
		} else {
			encode = append(encode, string(pwdMap[string(message[i])]))
		}
	}
	for i := 0; i < len(encode); i++ {
		trueEncode += encode[i]
	}
	return trueEncode
}

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key, message))
}
