package main

import "fmt"

func reverseVowels(s string) string {
	vowelInd, vowelSlice := make([]int, 0), make([]byte, 0)
	for i, v := range s {
		if v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' ||
			v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowelInd = append(vowelInd, i)
			vowelSlice = append(vowelSlice, byte(v))
		}
	}
	l := len(vowelSlice)
	for i := 0; i < l/2; i++ {
		vowelSlice[i], vowelSlice[l-1-i] = vowelSlice[l-1-i], vowelSlice[i]
	}
	str := []byte(s)
	for i := 0; i < l; i++ {
		str[vowelInd[i]] = vowelSlice[i]
	}
	return string(str)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("ae"))
}
