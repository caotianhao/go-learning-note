package main

import (
	"fmt"
	"strings"
)

func reverseWords151(s string) (ret string) {
	splitStr := strings.Split(s, " ")
	l := len(splitStr)
	for i := 0; i < l>>1; i++ {
		splitStr[i], splitStr[l-1-i] = splitStr[l-1-i], splitStr[i]
	}
	s1 := strings.TrimSpace(strings.Join(splitStr, " "))
	for i := range s1 {
		if s1[i] == ' ' && s1[i+1] == ' ' {
			ret += ""
		} else {
			ret += string(s1[i])
		}
	}
	return ret
}

func main() {
	fmt.Println(reverseWords151("the sky is blue"))
	fmt.Println(reverseWords151("           the               sky is blue    "))
	fmt.Println(reverseWords151("love you"))
	fmt.Println(reverseWords151("love"))
}
