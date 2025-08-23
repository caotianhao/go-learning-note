package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	rMap, mMap := map[byte]int{}, map[byte]int{}
	for _, v := range ransomNote {
		rMap[byte(v)]++
	}
	for _, v := range magazine {
		mMap[byte(v)]++
	}
	for i, v := range rMap {
		if mv, ok := mMap[i]; !(ok && mv >= v) {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(canConstruct("aa", "aab"))
	//fmt.Println(canConstruct("ba", "aab"))
	//fmt.Println(canConstruct("baa", "aab"))
	//fmt.Println(canConstruct("bba", "aab"))
	fmt.Println(canConstruct("sdfhgfjgfgsd", "aabsdfhgfjgfgsdfhf"))
}
