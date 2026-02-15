package main

import "fmt"

func lengthOfLastWord(s string) int {
	l, cnt, flag := len(s), 0, false
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if flag {
				break
			}
			continue
		} else {
			flag = true
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("       s            "))
	fmt.Println(lengthOfLastWord("s            "))
	fmt.Println(lengthOfLastWord("            s"))
	fmt.Println(lengthOfLastWord("a            s"))
}
