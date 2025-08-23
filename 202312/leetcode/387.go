package main

import (
	"fmt"
	"sort"
)

func firstUniqChar(s string) int {
	l := len(s)
	stringMap, retSlice := map[byte]int{}, make([]int, 0)
	for _, v := range s {
		stringMap[byte(v)]++
	}
	for i, v := range stringMap {
		if v != 1 {
			delete(stringMap, i)
		}
	}
	if len(stringMap) == 0 {
		return -1
	}
	for i := 0; i < l; i++ {
		if _, ok := stringMap[s[i]]; ok {
			retSlice = append(retSlice, i)
		}
	}
	sort.Ints(retSlice)
	return retSlice[0]
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("l"))
}
