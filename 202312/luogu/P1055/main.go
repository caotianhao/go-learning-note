package main

import (
	"fmt"
)

func main() {
	var c byte
	var code string
	for i := 0; i < 13; i++ {
		_, _ = fmt.Scanf("%c", &c)
		code += string(c)
	}
	sum := int((code[0] - 48) + 2*(code[2]-48) + 3*(code[3]-48) + 4*(code[4]-48) +
		5*(code[6]-48) + 6*(code[7]-48) + 7*(code[8]-48) + 8*(code[9]-48) + 9*(code[10]-48))
	if code[12] != 'X' && sum%11 == int(code[12]-48) || code[12] == 'X' && sum%11 == 10 {
		fmt.Printf("Right")
	} else {
		tmp := sum % 11
		if tmp != 10 {
			fmt.Printf("%s", code[:12]+string(byte(tmp+48)))
		} else {
			fmt.Printf("%s", code[:12]+"X")
		}
	}
}
