package main

import (
	"fmt"
	"reflect"
)

func isAnagram2032(s string, t string) bool {
	ms, mt := map[byte]int{}, map[byte]int{}
	for _, v := range s {
		ms[byte(v)]++
	}
	for _, v := range t {
		mt[byte(v)]++
	}
	return reflect.DeepEqual(ms, mt) && s != t
}

func main() {
	fmt.Println(isAnagram2032("a", "a"))
}
