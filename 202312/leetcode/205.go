package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	map205, tt, l := map[byte]byte{}, map[byte]byte{}, len(s)
	for i := 0; i < l; i++ {
		if v, ok := map205[s[i]]; ok && v != t[i] {
			return false
		}
		if _, ok := tt[t[i]]; ok && map205[s[i]] != t[i] {
			return false
		}
		map205[s[i]] = t[i]
		tt[t[i]] = t[i]
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("abc", "def"))   //t
	fmt.Println(isIsomorphic("abcc", "deff")) //t
	fmt.Println(isIsomorphic("abcd", "deff")) //f
	fmt.Println(isIsomorphic("abcc", "defk")) //f
	fmt.Println(isIsomorphic("badc", "baba")) //f
}
