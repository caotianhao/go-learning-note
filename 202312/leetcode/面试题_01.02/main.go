package main

import (
	"fmt"
	"reflect"
)

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1, m2 := map[byte]int{}, map[byte]int{}
	for _, v := range s1 {
		m1[byte(v)]++
	}
	for _, v := range s2 {
		m2[byte(v)]++
	}
	return reflect.DeepEqual(m1, m2)
}

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}
