package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	la, lb := len(a), len(b)
	if la > lb {
		bb := ""
		for i := 0; i < la-lb; i++ {
			bb += "0"
		}
		bb += b
		b = bb
	} else if lb > la {
		aa := ""
		for i := 0; i < lb-la; i++ {
			aa += "0"
		}
		aa += a
		a = aa
	}
	//fmt.Println(a, b)
	l, s := len(a), ""
	jw := false
	for i := l - 1; i >= 0; i-- {
		if a[i] == '0' && b[i] == '0' {
			if jw {
				s += "1"
				jw = false
			} else {
				s += "0"
				jw = false
			}
		} else if a[i] == '0' && b[i] == '1' {
			if jw {
				s += "0"
				jw = true
			} else {
				s += "1"
				jw = false
			}
		} else if a[i] == '1' && b[i] == '0' {
			if jw {
				s += "0"
				jw = true
			} else {
				s += "1"
				jw = false
			}
		} else {
			if jw {
				s += "1"
				jw = true
			} else {
				s += "0"
				jw = true
			}
		}
	}
	str := []rune(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	if (str[0] == '0') || (str[0] == '1' && jw == true) {
		return "1" + string(str)
	} else {
		return string(str)
	}
}

func main() {
	fmt.Println(addBinary("10010", "10"))
	fmt.Println(addBinary("0", "0"))
}
