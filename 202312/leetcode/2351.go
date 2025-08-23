package main

import "fmt"

func repeatedCharacter(s string) byte {
	l, charMap := len(s), make(map[string]int)
	var myByte byte
	for i := 0; i < l; i++ {
		charMap[string(s[i])]++
		if charMap[string(s[i])] == 2 {
			myByte = s[i]
			break
		}
	}
	return myByte
}

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}
