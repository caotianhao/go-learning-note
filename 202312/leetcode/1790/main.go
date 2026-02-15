package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	l1, l2, map1, map2 := len(s1), len(s2), make(map[string]int), make(map[string]int)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		map1[string(s1[i])]++
	}
	for i := 0; i < l2; i++ {
		map2[string(s2[i])]++
	}
	for i := 'a'; i <= 'z'; i++ {
		if map1[string(i)] != map2[string(i)] {
			return false
		}
	}
	notEqual := 0
	for i := 0; i < l1; i++ {
		if s1[i] != s2[i] {
			notEqual++
		}
	}
	if notEqual <= 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(areAlmostEqual("siyolsdcjthwsiplccjbuceoxmpjgrauocx", "siyolsdcjthwsiplccpbuceoxmjjgrauocx"))
}
