package main

import "fmt"

func similarPairs(words []string) (cnt int) {
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if judge2506(words[i], words[j]) {
				cnt++
			}
		}
	}
	return
}

func judge2506(a, b string) bool {
	ma, mb := map[byte]int{}, map[byte]int{}
	for _, v := range a {
		ma[byte(v)]++
	}
	for _, v := range b {
		mb[byte(v)]++
	}
	flag1, flag2 := true, true
	for i := range mb {
		if _, ok := ma[i]; !ok {
			flag1 = false
			break
		}
	}
	for i := range ma {
		if _, ok := mb[i]; !ok {
			flag2 = false
			break
		}
	}
	return flag1 && flag2
}

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
	fmt.Println(similarPairs([]string{"aabb", "ab", "ba"}))
	fmt.Println(similarPairs([]string{"aabb"}))
}
