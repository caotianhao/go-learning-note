package main

import "fmt"

type WordsFrequency struct {
	hm map[string]int
}

func Constructor1602(book []string) WordsFrequency {
	help := map[string]int{}
	for _, v := range book {
		help[v]++
	}
	return WordsFrequency{hm: help}
}

func (wf *WordsFrequency) Get(word string) int {
	return wf.hm[word]
}

func main() {
	c := Constructor1602([]string{"i", "have", "an", "apple", "he", "have", "a", "pen"})
	fmt.Println(c.Get("have"))
}
