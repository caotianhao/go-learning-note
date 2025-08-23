package main

import "fmt"

func longestNiceSubstring(s string) string {
	l := len(s)
	want := ""
	if isNice(s) {
		return s
	}
	if l <= 2 {
		return ""
	}
	for i := l - 1; i >= 2; i-- {
		for j := 0; j < l-i; j++ {
			if isNice(s[j : j+i]) {
				return s[j : j+i]
			}
		}
		if isNice(s[l-i:]) {
			want = s[l-i:]
			break
		}
	}
	return want
}

func isNice(s string) bool {
	big := make(map[byte]int)
	small := make(map[byte]int)
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			small[byte(v)] = int(v)
		}
		if v >= 'A' && v <= 'Z' {
			big[byte(v)] = int(v)
		}
	}
	//fmt.Println(small, big)
	if len(big) != len(small) {
		return false
	} else {
		flag := true
		for _, v := range big {
			if _, ok := small[byte(v)+32]; !ok {
				flag = false
			}
		}
		return flag
	}
}

func main() {
	//fmt.Println(isNice("aAbbBB"))
	fmt.Println(longestNiceSubstring("YazaAay"))
	fmt.Println(longestNiceSubstring("dDzeE"))
}
