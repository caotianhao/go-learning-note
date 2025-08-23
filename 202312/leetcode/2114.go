package main

import "fmt"

func mostWordsFound(sentences []string) int {
	maxWF := 0
	for i := 0; i < len(sentences); i++ {
		if howManyWords(sentences[i]) > maxWF {
			maxWF = howManyWords(sentences[i])
		}
	}
	return maxWF
}

func howManyWords(s string) int {
	space := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			space++
		}
	}
	return space + 1
}

func main() {
	arr := []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(arr))
}
