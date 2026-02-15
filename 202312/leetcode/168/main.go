package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ret := ""
	for _, v := range decTo26168(columnNumber) {
		ret += string(rune(v + 64))
	}
	return ret
}

func decTo26168(n int) (slice168 []int) {
	for n != 0 {
		if n%26 != 0 {
			slice168 = append(slice168, n%26)
			n /= 26
		} else {
			slice168 = append(slice168, 26)
			n = n/26 - 1
		}
	}
	return reverse168(slice168)
}

func reverse168(a []int) []int {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
	return a
}

func main() {
	fmt.Println(convertToTitle(45342436))
	fmt.Println(convertToTitle(2147483647))
	fmt.Println(convertToTitle(18278))
}
