package main

import "fmt"

func maximumOddBinaryNumber(s string) (res string) {
	cnt, n := 0, len(s)
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}
	for i := 0; i < cnt-1; i++ {
		res += "1"
	}
	for i := 0; i < n-cnt; i++ {
		res += "0"
	}
	return res + "1"
}

func main() {
	fmt.Println(maximumOddBinaryNumber("0100010101"))
	fmt.Println(maximumOddBinaryNumber("1110000001"))
}
