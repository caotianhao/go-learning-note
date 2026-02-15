package main

import "fmt"

//func lastSubstring(s string) (res string) {
//	map1163, maxLetter := map[rune]struct{}{}, ' '
//	for _, v := range s {
//		map1163[v] = struct{}{}
//	}
//	for i := 'z'; i >= 'a'; i-- {
//		if _, ok := map1163[i]; ok {
//			maxLetter = i
//			break
//		}
//	}
//	for i, v := range s {
//		if v == maxLetter {
//			tmp := s[i:]
//			if tmp > res {
//				res = tmp
//			}
//		}
//	}
//	return
//}

func lastSubstring(s string) string {
	l := len(s)
	i, j, k := 0, 1, 0
	for j+k < l {
		if s[i+k] == s[j+k] {
			k++
		} else if s[i+k] < s[j+k] {
			i += k + 1
			k = 0
			if i >= j {
				j = i + 1
			}
		} else {
			j += k + 1
			k = 0
		}
	}
	return s[i:]
}

func main() {
	fmt.Println(lastSubstring("leetcode"))
	fmt.Println(lastSubstring("fdghssbkshifuhkjvbgiygfkjvbdfhksgfahdfsrgshssdh"))
}
