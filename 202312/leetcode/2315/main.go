package main

import "fmt"

func countAsterisks(s string) int {
	cntLine, cntStar, l := 0, 0, len(s)
	for i := 0; i < l; i++ {
		if s[i] == '|' {
			cntLine++
		}
		if s[i] == '*' && cntLine%2 == 0 {
			cntStar++
		}
	}
	return cntStar
}

func main() {
	s := "****************"
	fmt.Println(countAsterisks(s))
}
