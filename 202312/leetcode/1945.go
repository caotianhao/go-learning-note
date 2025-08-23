package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	strInt, sum := "", 0
	for i := 0; i < len(s); i++ {
		strInt += strconv.FormatInt(int64(s[i]-96), 10)
	}
	for i := 0; i < len(strInt); i++ {
		sum += int(strInt[i] - 48)
	}
	for i := 0; i < k-1; i++ {
		sum = getSum1945(sum)
	}
	return sum
}

func getSum1945(n int) (sum int) {
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return
}

func main() {
	fmt.Println(getLucky("zbax", 2))         //8
	fmt.Println(getLucky("fleyctuuajsr", 5)) //8
	fmt.Println(getLucky("fleycfdgjkdtssjtuuajsr", 13))
}
