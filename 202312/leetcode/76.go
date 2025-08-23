package main

import (
	"fmt"
	"sort"
)

//耗时很多
func minWindow(s string, t string) string {
	p, q, retSlice, ret := 0, 0, make([]string, 0), ""
	mapS, mapT := map[rune]int{}, map[rune]int{}
	for _, v := range t {
		mapT[v]++
	}
	for q < len(s) {
		mapS[rune(s[q])]++
		if sIsHaveT(mapS, mapT) {
			retSlice = append(retSlice, s[p:q+1])
		}
		for sIsHaveT(mapS, mapT) {
			mapS[rune(s[p])]--
			retSlice = append(retSlice, s[p:q+1])
			p++
		}
		q++
	}
	sort.Slice(retSlice, func(i, j int) bool {
		return len(retSlice[i]) < len(retSlice[j])
	})
	if len(retSlice) != 0 {
		ret = retSlice[0]
	}
	return ret
}

func sIsHaveT(s, t map[rune]int) bool {
	for i, v := range s {
		if vt, _ := t[i]; vt > v {
			return false
		}
	}
	for i := range t {
		if _, ok := s[i]; !ok {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
	fmt.Println(minWindow("a", "aa"))
}
