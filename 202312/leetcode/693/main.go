package main

import "fmt"

func hasAlternatingBits(n int) bool {
	l := len(dec2Bin693(n))
	for i := 0; i < l; i++ {
		if (i&1 == 0 && dec2Bin693(n)[i] == 1-l&1) || (i&1 == 1 && dec2Bin693(n)[i] == l&1) {
			return false
		}
	}
	return true
}

func dec2Bin693(n int) (slice693 []int) {
	for n != 0 {
		slice693 = append(slice693, n&1)
		n /= 2
	}
	return
}

func main() {
	fmt.Println(hasAlternatingBits(94))
	fmt.Println(hasAlternatingBits(1))
	fmt.Println(hasAlternatingBits(2))
	fmt.Println(hasAlternatingBits(3))
}
