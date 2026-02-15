package main

import "fmt"

func stringMatching(words []string) []string {
	l, ret := len(words), make([]string, 0)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i != j {
				if isSubstr1408(words[i], words[j]) {
					ret = append(ret, words[j])
				}
			}
		}
	}
	return deleteRedundant1408(ret)
}

func isSubstr1408(s, s1 string) bool {
	ls, ls1 := len(s), len(s1)
	if ls1 > ls {
		return false
	} else if ls1 == ls {
		return s == s1
	} else {
		for i := 0; i < ls-ls1+1; i++ {
			if s[i:i+ls1] == s1 {
				return true
			}
		}
		return false
	}
}

func deleteRedundant1408(s []string) []string {
	map1408, slice1408 := map[string]string{}, make([]string, 0)
	for _, v := range s {
		map1408[v] = v
	}
	for _, v := range map1408 {
		slice1408 = append(slice1408, v)
	}
	return slice1408
}

func main() {
	//fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	fmt.Println(stringMatching([]string{"leetcode", "et", "code"}))
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
	//fmt.Println(stringMatching([]string{"blue", "green", "bu"}))
	//fmt.Println(isSubstr1408("mass", "as"))
}
