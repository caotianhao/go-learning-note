package main

import "fmt"

func removeVowels(s string) string {
	res := ""
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			continue
		}
		res += string(v)
	}
	return res
}

func main() {
	fmt.Println(removeVowels("afgd"))
}
