package main

import (
	"fmt"
)

// s2 是否包含 s1 的任何一种排列
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	//m1 := map[byte]int{}
	//for _, v := range s1 {
	//	m1[byte(v)]++
	//}
	arr1 := [26]byte{}
	//arr1 := make([]byte, 26)
	for _, v := range s1 {
		arr1[byte(v-'a')]++
	}
	for i := 0; i+l1 <= l2; i++ {
		//-------------------------------------------------
		// 要是用 map 判断的话就很耗时，用 slice 就好多了
		//-------------------------------------------------
		//m2 := map[byte]int{}
		//for _, v := range s2[i : i+l1] {
		//	m2[byte(v)]++
		//}
		arr2 := [26]byte{}
		//arr2 := make([]byte, 26)
		for _, v := range s2[i : i+l1] {
			arr2[byte(v-'a')]++
		}
		//if reflect.DeepEqual(m1, m2) {
		//	return true
		//}
		if arr1 == arr2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
	fmt.Println(checkInclusion("ab", "eidboaoo")) // false
}
