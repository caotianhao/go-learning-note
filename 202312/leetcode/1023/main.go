package main

import (
	"fmt"
	"strings"
)

func camelMatch(queries []string, pattern string) (res []bool) {
	pat := []byte(pattern)
	var check func(str, pat []byte) bool
	check = func(str, pat []byte) bool {
		cnt, j := 0, 0
		for i := 0; i < len(pat); {
			flag := false
			for ; j < len(str); j++ {
				if str[j] == pat[i] {
					i++
					if str[j] >= 'A' && str[j] <= 'Z' {
						str[j] += 32
					}
					cnt++
					flag = true
					break
				}
			}
			if !flag {
				i++
			}
		}
		return cnt == len(pat) && strings.ToLower(string(str)) == string(str)
	}
	for _, v := range queries {
		res = append(res, check([]byte(v), pat))
	}
	return
}

func main() {
	//s := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	//fmt.Println(camelMatch(s, "FB"))                                                                   // tfttf
	//fmt.Println(camelMatch(s, "FoBa"))                                                                 // tftff
	//fmt.Println(camelMatch(s, "FoBaT"))                                                                // ftfff
	//fmt.Println(camelMatch([]string{"A", "a"}, "a"))                                                   // ft
	fmt.Println(camelMatch([]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "CooP")) // fft
}
