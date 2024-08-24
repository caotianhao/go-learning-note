package main

import "fmt"

//func countOPPO(s string) int {
//	n := len(s)
//	if n < 4 {
//		return 0
//	}
//	left, right, count, v := 0, 3, 0, 0
//	for right < n {
//		if isVowel(s[right]) {
//			v++
//		}
//		if right-left+1 > 4 {
//			if isVowel(s[left]) {
//				v--
//			}
//			left++
//		}
//		if right-left+1 == 4 && v == 2 {
//			if isOPPO(s[left : right+1]) {
//				count++
//			}
//		}
//		right++
//	}
//	return count
//}

func countOPPO(s string) (cnt int) {
	n := len(s)
	for i := 0; i < n; i++ {
		if !isVowel(s[i]) {
			continue
		}
		for j := i + 1; j < n; j++ {
			if isVowel(s[j]) {
				continue
			}
			for k := j + 1; k < n; k++ {
				if isVowel(s[k]) || s[k] != s[j] {
					continue
				}
				for m := k + 1; m < n; m++ {
					if !isVowel(s[m]) || s[m] != s[i] {
						continue
					}
					if isOPPO(string(s[i]) + string(s[j]) + string(s[k]) + string(s[m])) {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func isOPPO(s string) bool {
	var f1, f2 bool
	if isVowel(s[0]) && isVowel(s[3]) && s[0] == s[3] {
		f1 = true
	}
	if !isVowel(s[1]) && !isVowel(s[2]) && s[1] == s[2] {
		f2 = true
	}
	return f1 && f2
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func main() {
	s := ""
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%d", countOPPO(s))
}
