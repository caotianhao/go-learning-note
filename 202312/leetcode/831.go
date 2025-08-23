package main

import (
	"fmt"
	"strings"
)

func maskPII(s string) string {
	s = strings.ToLower(s)
	for i, v := range s {
		if v == '@' {
			return mail831(s, i)
		}
	}
	return phone831(s)
}

func mail831(s string, atInd int) string {
	name, domain := s[:atInd], s[atInd+1:]
	return string(name[0]) + "*****" + string(name[atInd-1]) + "@" + domain
}

func phone831(s string) string {
	number := ""
	for _, v := range s {
		if v >= '0' && v <= '9' {
			number += string(v)
		}
	}
	l := len(number)
	lastFour := number[l-4:]
	if l == 10 {
		return "***-***-" + lastFour
	} else if l == 11 {
		return "+*-***-***-" + lastFour
	} else if l == 12 {
		return "+**-***-***-" + lastFour
	}
	return "+***-***-***-" + lastFour
}

func main() {
	fmt.Println(maskPII("LeetCode@LeetCode.com"))
	fmt.Println(maskPII("86-(10)12345678"))
}
