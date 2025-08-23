package main

import "fmt"

func strongPasswordCheckerII(password string) bool {
	//它有至少 8个字符。
	//至少包含 一个小写英文字母。
	//至少包含 一个大写英文字母。
	//至少包含 一个数字。
	//至少包含 一个特殊字符。特殊字符为："!@#$%^&*()-+"中的一个。
	//它不包含2个连续相同的字符（比方说"aab"不符合该条件，但是"aba"符合该条件）。
	l := len(password)
	if l < 8 {
		return false
	}
	for i := 0; i < l-1; i++ {
		if password[i] == password[i+1] {
			return false
		}
	}
	haveSmall, haveBig, haveNum, haveSpec := false, false, false, false
	for _, v := range password {
		if v >= 'a' && v <= 'z' {
			haveSmall = true
		}
		if v >= 'A' && v <= 'Z' {
			haveBig = true
		}
		if v >= '0' && v <= '9' {
			haveNum = true
		}
		if v == '!' || v == '@' || v == '#' || v == '$' || v == '%' ||
			v == '^' || v == '&' || v == '*' || v == '(' || v == ')' ||
			v == '-' || v == '+' {
			haveSpec = true
		}
	}
	return haveSpec && haveBig && haveSmall && haveNum
}

func main() {
	fmt.Println(strongPasswordCheckerII("IloveLe3tcode!"))
}
