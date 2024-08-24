package main

import (
	"fmt"
	"strings"
)

func isPhoneNumber(phoneNumber string) bool {
	phoneNumber = strings.Trim(phoneNumber, " )-/*@#$%^&*")
	b := strings.Index(phoneNumber, ")")
	if b == -1 {
		cnt := 0
		notNum := make([]int, 0)
		notNumF := make([]rune, 0)
		for i, v := range phoneNumber {
			if !isNumber(v) {
				notNum = append(notNum, i)
				notNumF = append(notNumF, v)
			} else {
				cnt++
			}
		}
		if cnt == 11 {
			if len(notNum) == 0 {
				return true
			} else if len(notNum) == 1 {
				return notNum[0] == 3
			} else if len(notNum) == 2 {
				return notNum[0] == 3 && notNum[1] == 8 && notNumF[0] == notNumF[1]
			}
		}
		return false
	} else {
		return isPhoneNumber(phoneNumber[b+1:])
	}
}

func isNumber(b rune) bool {
	return b >= '0' && b <= '9'
}

func main() {
	fmt.Println(isPhoneNumber("123-45678900"))
}
