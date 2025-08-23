package main

import "fmt"

func finalString(s string) string {
	res := ""
	for _, v := range s {
		if v == 'i' {
			res = reverse2810(res)
		} else {
			res += string(v)
		}
	}
	return res
}

func reverse2810(s string) string {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		r[i], r[l-1-i] = r[l-i-1], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(finalString("pii"))
}
