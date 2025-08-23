package main

import "fmt"

func uncommonFromSentences(s1 string, s2 string) []string {
	var slice1, slice2, ret []string
	index1, index2 := 0, 0
	map1, map2 := map[string]int{}, map[string]int{}
	for i, v := range s1 {
		if v == ' ' {
			slice1 = append(slice1, s1[index1:i])
			index1 = i + 1
		}
	}
	slice1 = append(slice1, s1[index1:])
	for i, v := range s2 {
		if v == ' ' {
			slice2 = append(slice2, s2[index2:i])
			index2 = i + 1
		}
	}
	slice2 = append(slice2, s2[index2:])
	for _, v := range slice1 {
		map1[v]++
	}
	for _, v := range slice2 {
		map2[v]++
	}
	for i, v := range map1 {
		if _, ok := map2[i]; !ok && v == 1 {
			ret = append(ret, i)
		}
	}
	for i, v := range map2 {
		if _, ok := map1[i]; !ok && v == 1 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
