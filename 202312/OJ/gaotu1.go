package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type wordFrequency struct {
	word string
	freq int
}

func sortWordFreq(w map[string]int) (r []wordFrequency) {
	for word, freq := range w {
		r = append(r, wordFrequency{word, freq})
	}
	sort.Slice(r, func(i, j int) bool {
		if r[i].freq == r[j].freq {
			return r[i].word < r[j].word
		}
		return r[i].freq > r[j].freq
	})
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words := strings.Split(s, " ")
	//fmt.Println(words)
	wordFreq := make(map[string]int)
	for _, w := range words {
		wordFreq[w]++
	}
	r := sortWordFreq(wordFreq)
	for _, pair := range r {
		fmt.Println(pair.word, pair.freq)
	}
}
