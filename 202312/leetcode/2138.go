package main

import "fmt"

func divideString(s string, k int, fill byte) (slice2138 []string) {
	l := len(s)
	for i := 0; i < l/k; i++ {
		slice2138 = append(slice2138, s[i*k:(i+1)*k])
	}
	if l%k == 0 {
		return
	}
	tmp := s[l/k*k:]
	for i := 0; i < k-(l-l/k*k); i++ {
		tmp += string(fill)
	}
	slice2138 = append(slice2138, tmp)
	return
}

func main() {
	fmt.Println(divideString("abcdefghiq", 3, 'x'))
	fmt.Println(divideString("ab", 32, 'x'))
}
