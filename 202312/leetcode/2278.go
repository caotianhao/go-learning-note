package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	l, cnt := len(s), 0
	for i := 0; i < l; i++ {
		if s[i] == letter {
			cnt++
		}
	}
	return 100 * cnt / l
}

func main() {
	s := "football"
	var letter byte = 'o'
	fmt.Println(percentageLetter(s, letter))
}
