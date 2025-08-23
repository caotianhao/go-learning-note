package main

import "fmt"

func wordPattern(pattern string, s string) bool {
	lp, ls := len(pattern), len(verse290(s))
	if lp != ls {
		return false
	}
	map290, mv := map[byte]string{}, map[string]string{}
	for i := 0; i < lp; i++ {
		if v, ok := map290[pattern[i]]; ok && v != verse290(s)[i] {
			return false
		}
		if _, ok2 := mv[verse290(s)[i]]; ok2 && map290[pattern[i]] != verse290(s)[i] {
			return false
		}
		map290[pattern[i]] = verse290(s)[i]
		mv[verse290(s)[i]] = verse290(s)[i]
	}
	return true
}

func verse290(s string) (slice290 []string) {
	tmp := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			slice290 = append(slice290, tmp)
			tmp = ""
			continue
		}
		tmp += string(s[i])
	}
	slice290 = append(slice290, tmp)
	return
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog")) //t
	fmt.Println(wordPattern("abca", "dog cat cat dog")) //f
	fmt.Println(wordPattern("abcc", "dog cat cat dog")) //f
	fmt.Println(wordPattern("abcd", "dog cat szz gfb")) //f
	//fmt.Println(verse290("dog cat cat dog"))            //t
}
