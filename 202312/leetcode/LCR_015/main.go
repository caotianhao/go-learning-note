package main

import "fmt"

func findAnagrams2015(s string, p string) []int {
	ls, lp, ret := len(s), len(p), make([]int, 0)
	if ls < lp {
		return ret
	}
	var sCnt, pCnt [26]int
	for i, v := range p {
		sCnt[s[i]-'a']++
		pCnt[v-'a']++
	}
	if sCnt == pCnt {
		ret = append(ret, 0)
	}
	for i := 0; i < ls-lp; i++ {
		sCnt[s[i]-'a']--
		sCnt[s[i+lp]-'a']++
		if sCnt == pCnt {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	fmt.Println(findAnagrams2015("abab", "ab"))
}
