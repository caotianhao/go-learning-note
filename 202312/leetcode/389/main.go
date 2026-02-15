package main

import "fmt"

func findTheDifference(s string, t string) byte {
	ms, mt := map[byte]int{}, map[byte]int{}
	for _, v := range s {
		ms[byte(v)]++
	}
	for _, v := range t {
		mt[byte(v)]++
	}
	for i := range mt {
		if mt[i] != ms[i] {
			return i
		}
	}
	return ' '
}

func main() {
	fmt.Println(findTheDifference("abdasfngfdshhf", "ngfdssabdasfhhf"))
	fmt.Println(findTheDifference("", "b"))
}
