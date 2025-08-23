package main

import "fmt"

func freqAlphabets(s string) string {
	str, l := "", len(s)
	for i := 0; i < l-2; i++ {
		if s[i] == '#' || s[i+1] == '#' {
			continue
		}
		if s[i+2] != '#' {
			str += string(s[i] + 48)
		} else {
			str += string(byte(int(s[i]-48)*10 + int(s[i+1]-48) + 96))
		}
	}
	if s[l-2] != '#' && s[l-1] != '#' {
		str += string(s[l-2] + 48)
		str += string(s[l-1] + 48)
	}
	if s[l-2] == '#' && s[l-1] != '#' {
		str += string(s[l-1] + 48)
	}
	return str
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}
