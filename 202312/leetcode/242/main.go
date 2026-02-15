package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms, mt := map[byte]int{}, map[byte]int{}
	for _, v := range s {
		ms[byte(v)]++
	}
	for _, v := range t {
		mt[byte(v)]++
	}
	return reflect.DeepEqual(ms, mt)
}

func main() {
	fmt.Println(isAnagram("abc", "bac"))
}
