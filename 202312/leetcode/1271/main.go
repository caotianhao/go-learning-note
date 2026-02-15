package main

import (
	"fmt"
	"strconv"
)

func toHexspeak(num string) string {
	h, _ := strconv.ParseInt(num, 10, 64)
	res := fmt.Sprintf("%X", h)
	ans := ""
	for _, v := range res {
		if v == '2' || v == '3' || v == '4' || v == '5' ||
			v == '6' || v == '7' || v == '8' || v == '9' {
			return "ERROR"
		} else if v == '0' {
			ans += "O"
		} else if v == '1' {
			ans += "I"
		} else {
			ans += string(v)
		}
	}
	return ans
}

func main() {
	fmt.Println(toHexspeak("257"))
	fmt.Println(toHexspeak("22"))
	fmt.Println(toHexspeak("747823223228"))
}
