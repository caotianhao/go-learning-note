package main

import "fmt"

func checkZeroOnes(s string) bool {
	l, zero, one, maxZero, maxOne := len(s), 0, 0, 0, 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			one++
			zero = 0
			if one > maxOne {
				maxOne = one
			}
		} else {
			zero++
			one = 0
			if zero > maxZero {
				maxZero = zero
			}
		}
	}
	fmt.Println(maxOne, maxZero)
	return maxOne-maxZero > 0
}

func main() {
	fmt.Println(checkZeroOnes("1110000110001010000101000110011001"))
	fmt.Println(checkZeroOnes("1"))
	fmt.Println(checkZeroOnes("0"))
}
