package main

import "fmt"

//func reverseWords(s string) string {
//	l, str := len(s), ""
//	index, space, haveSpace := 0, 0, false
//	for i := 0; i < l; i++ {
//		index++
//		if s[i] == ' ' {
//			str += reverse557(s[i-index+1 : i])
//			str += " "
//			index, space, haveSpace = 0, i, true
//		}
//		if i == l-1 && haveSpace {
//			str += reverse557(s[space+1:])
//		}
//	}
//	if !haveSpace {
//		str += reverse557(s)
//	}
//	return str
//}
//
//func reverse557(s string) string {
//	l := len(s)
//	str := []rune(s)
//	for i := 0; i < l/2; i++ {
//		str[i], str[l-1-i] = str[l-1-i], str[i]
//	}
//	return string(str)
//}

func reverseWords(s string) (str string) {
	l, r, flag := 0, 0, true
	for r < len(s) {
		if s[r] == ' ' {
			flag = false
			str += reverse557(s[l:r])
			str += " "
			l = r + 1
		}
		r++
	}
	if flag {
		return reverse557(s)
	} else {
		str += reverse557(s[l:])
		return str
	}
}

func reverse557(s string) string {
	str := []byte(s)
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("Le"))
	//fmt.Println(reverse557("Let's"))
}
