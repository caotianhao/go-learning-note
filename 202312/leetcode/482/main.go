package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) (str string) {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1))
	l := len(s)
	if first := l % k; first == 0 {
		for j := 0; j < l/k; j++ {
			for i := 0; i < k; i++ {
				str += string(s[i+j*k])
			}
			str += "-"
		}
	} else {
		for i := 0; i < first; i++ {
			str += string(s[i])
		}
		str += "-"
		for j := 0; j < l/k; j++ {
			for i := 0; i < k; i++ {
				str += string(s[i+j*k+first])
			}
			str += "-"
		}
	}
	if len(str) > 0 {
		return str[:len(str)-1]
	}
	return str
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w-a", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 3))
	fmt.Println(licenseKeyFormatting("2", 1))
	fmt.Println(licenseKeyFormatting("---", 3))
}
