package main

import "fmt"

func removeOuterParentheses(s string) string {
	left, right, l, temp, tempSlice := 0, 0, len(s), "", make([]string, 0)
	for i := 0; i < l; i++ {
		temp = ""
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			temp = s[1+i-left-right : i+1]
			left, right = 0, 0
			//fmt.Println(temp)
			tempSlice = append(tempSlice, temp)
			//fmt.Println("aa", tempSlice)
		}
	}
	//fmt.Println(tempSlice)
	temp = ""
	for i := range tempSlice {
		if len(tempSlice[i]) == 2 {
			temp += ""
		} else {
			temp += tempSlice[i][1 : len(tempSlice[i])-1]
		}
	}
	return temp
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
