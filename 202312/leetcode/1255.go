package main

import (
	"fmt"
	"math"
)

func maxScoreWords(words []string, letters []byte, score []int) (maxN int) {
	allLetter, l := [26]byte{}, len(words)
	for _, v := range letters {
		allLetter[v-97]++
	}
	wordScore := make([]int, l)
	for i, v := range words {
		if !judge1255(v, allLetter) {
			wordScore[i] = math.MinInt64
		} else {
			for j := 0; j < len(v); j++ {
				wordScore[i] += score[v[j]-97]
			}
		}
	}
	for i := 0; i < 1<<l; i++ {
		s, num := "", 0
		for j := 0; j < i; j++ {
			if i>>j&1 == 1 {
				s += words[l-j-1]
				num += wordScore[l-j-1]
			}
		}
		if judge1255(s, allLetter) {
			maxN = max(maxN, num)
		}
	}
	return
}

func judge1255(str string, m [26]byte) bool {
	mm := [26]byte{}
	for _, v := range str {
		mm[v-97]++
	}
	for i, v := range mm {
		if v > m[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(maxScoreWords([]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5,
			0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
