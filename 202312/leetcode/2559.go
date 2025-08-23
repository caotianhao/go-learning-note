package main

import "fmt"

func vowelStrings2559(words []string, queries [][]int) (res []int) {
	l := len(words)
	isVowel := make([]int, l)
	for i := 0; i < l; i++ {
		if judge2559(words[i]) {
			isVowel[i] = 1
		}
	}
	sum := make([]int, l+1)
	for i := 1; i < l+1; i++ {
		sum[i] += sum[i-1] + isVowel[i-1]
	}
	for _, v := range queries {
		res = append(res, sum[v[1]+1]-sum[v[0]])
	}
	return
}

func judge2559(s string) bool {
	var isVowel func(b byte) bool
	isVowel = func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
	}
	return isVowel(s[0]) && isVowel(s[len(s)-1])
}

func main() {
	fmt.Println(vowelStrings2559([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
}
