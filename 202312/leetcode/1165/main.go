package main

import "fmt"

func calculateTime(keyboard, word string) (ans int) {
	wordInt := make([]int, 1)
	wordInt[0] = 0
	keyMap := map[byte]int{}
	for i := 0; i < 26; i++ {
		keyMap[keyboard[i]] = i
	}
	for i := 0; i < len(word); i++ {
		wordInt = append(wordInt, keyMap[word[i]])
	}
	for i := 0; i < len(wordInt)-1; i++ {
		ans += abs1165(wordInt[i], wordInt[i+1])
	}
	return
}

func abs1165(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode"))
}
