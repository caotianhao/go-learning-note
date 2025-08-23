package main

import (
	"fmt"
	"strconv"
)

func baseNeg2(n int) (res string) {
	if n == 0 {
		return "0"
	}
	for n != 0 {
		if n > 0 {
			res += strconv.Itoa(n % (-2))
			n /= -2
		} else if n < 0 {
			res += strconv.Itoa(-n % 2)
			n = n/(-2) + (-n % 2)
		}
	}
	r, l := []rune(res), len(res)
	for i := 0; i < l/2; i++ {
		r[i], r[l-1-i] = r[l-1-i], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(baseNeg2(0))  //0
	fmt.Println(baseNeg2(1))  //1
	fmt.Println(baseNeg2(2))  //110
	fmt.Println(baseNeg2(19)) //10111
	fmt.Println(baseNeg2(4))  //100
}
