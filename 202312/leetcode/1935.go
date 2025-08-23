package main

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {
	cnt := 0
	for _, v := range divide1935(text) {
		if isIn1935(v, brokenLetters) {
			cnt++
		}
	}
	return cnt
}

func isIn1935(str, str1 string) bool {
	m1935 := map[byte]int{}
	for _, v := range str1 {
		m1935[byte(v)]++
	}
	for _, v := range str {
		if _, ok := m1935[byte(v)]; ok {
			return false
		}
	}
	return true
}

func divide1935(s string) []string {
	str, slice1935 := "", make([]string, 0)
	for _, v := range s {
		if v == ' ' {
			slice1935 = append(slice1935, str)
			str = ""
			continue
		}
		str += string(v)
	}
	slice1935 = append(slice1935, str)
	return slice1935
}

func main() {
	fmt.Println(canBeTypedWords("hello world dda", "ad"))
	fmt.Println(canBeTypedWords("hello", "l"))
}
