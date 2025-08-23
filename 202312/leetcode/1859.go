package main

import "fmt"

func sortSentence(s string) string {
	l, str, targetStr := len(s), []rune(" "+s), ""
	spaceId := 0
	for h := 1; h <= 9; h++ {
		for i := 0; i < l+1; i++ {
			if int(str[i]) == h+48 {
				for j := i - 1; j >= 0; j-- {
					if str[j] == ' ' {
						spaceId = j
						break
					}
				}
				targetStr += string(str[spaceId+1 : i])
				targetStr += " "
			}
		}
	}
	return targetStr[:len(targetStr)-1]
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}
