package main

import "fmt"

func largestOddNumber(num string) string {
	haveOdd := false
	for _, v := range num {
		if v == '1' || v == '3' || v == '5' || v == '7' || v == '9' {
			haveOdd = true
		}
	}
	if !haveOdd {
		return ""
	} else {
		for i := len(num) - 1; i >= 0; i-- {
			if num[i] == '1' || num[i] == '3' || num[i] == '5' ||
				num[i] == '7' || num[i] == '9' {
				return num[:i+1]
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(largestOddNumber("165413263543"))
}
