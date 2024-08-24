package main

import "fmt"

func judgeString(s string) int {
	l := len(s)
	z, o := "0", "1"
	s1, s2 := "", ""
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			s1 += z
			s2 += o
		} else {
			s1 += o
			s2 += z
		}
	}
	a, b := 0, 0
	for i := 0; i < l; i++ {
		if s[i] != s1[i] {
			a++
		}
		if s[i] != s2[i] {
			b++
		}
	}
	return min(a, b)
}

func hard3(s string) (k int) {
	l := len(s)
	memo := map[string]int{}
	for i := 1; i < l; i++ {
		for j := 0; j < l; j++ {
			if j+i <= l {
				ss := s[j : j+i]
				if _, ok := memo[ss]; !ok {
					memo[ss] = judgeString(ss)
					k += memo[ss]
				} else {
					k += memo[ss]
				}
			}
		}
	}
	return
}

func main() {
	s := ""
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%d\n", hard3(s)+judgeString(s))
}
