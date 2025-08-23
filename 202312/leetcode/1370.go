package main

import "fmt"

func sortString(s string) string {
	l, charMap, res := len(s), make(map[string]int), ""
	for i := 0; i < l; i++ {
		charMap[string(s[i])]++
	}
	for i := 97; i < 123; i++ {
		for j := 0; j < l; j++ {
			if int(s[j]) == i {
				res += string(s[j])
				charMap[string(s[j])]--
				s = minusChar(s, s[j])
				l--
				break
			}
		}
	}
	if l != 0 {
		for i := 122; i > 96; i-- {
			for j := 0; j < l; j++ {
				if int(s[j]) == i && charMap[string(s[j])] != 0 {
					res += string(s[j])
					charMap[string(s[j])]--
					s = minusChar(s, s[j])
					l--
					break
				}
			}
		}
	}
	if len(s) != 0 {
		return res + sortString(s)
	}
	return res
}

func minusChar(s string, c byte) string {
	l, s1 := len(s), ""
LOOP:
	for i := 0; i < l; i++ {
		if s[i] == c {
			for j := 0; j < l; j++ {
				if j == i {
					continue
				}
				s1 += string(s[j])
			}
			break LOOP
		}
	}
	return s1
}

func main() {
	fmt.Println(sortString("leetcode"))
}
