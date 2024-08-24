package main

import "fmt"

func getWords(words []string) int {
	n := len(words)
	if n == 1 {
		return 1
	}
	hashMap := make(map[string]struct{})
	cntMap := make(map[string]int)
	hashMap[words[0]] = struct{}{}
	cnt := 1
	for i := 1; i < n; i++ {
		if _, ok := hashMap[words[i]]; !ok {
			cntMap[words[i]]++
			if cntMap[words[i]] == len(hashMap)+1 {
				hashMap[words[i]] = struct{}{}
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	n := 0
	strs := make([]string, 0)
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		s := ""
		_, _ = fmt.Scanf("%s", &s)
		strs = append(strs, s)
	}
	fmt.Println(getWords(strs))
}
