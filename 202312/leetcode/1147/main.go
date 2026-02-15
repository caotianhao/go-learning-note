package main

import "fmt"

//func longestDecomposition(text string) (cnt int) {
//	l := len(text)
//	left, right := "", ""
//	head, tail := 0, l-1
//	//fmt.Println("l", l)
//	for head < tail {
//		left += string(text[head])
//		right += string(text[tail])
//		head++
//		tail--
//		if left == reverse1147(right) {
//			cnt += 2
//			left, right = "", ""
//		}
//		//fmt.Println(left, reverse1147(right))
//	}
//	if left == "" && right == "" && l%2 == 0 {
//		return
//	}
//	return cnt + 1
//}

//func reverse1147(str string) string {
//	s := []byte(str)
//	l := len(s)
//	for i := 0; i < l/2; i++ {
//		s[i], s[l-1-i] = s[l-i-1], s[i]
//	}
//	return string(s)
//}

func longestDecomposition(text string) int {
	l := len(text)
	// 如果最后删着删着空了，那就是0
	if l == 0 {
		return 0
	}
	// 先判断第一个和最后一个是不是相等，然后前两个和后两个
	// 为啥是 <= 而不是 < 用aa试一下就知道了
	// aa的len为2，假设是<，那么直接return 1 了，显然不对
	for head := 1; head <= l/2; head++ {
		if text[:head] == text[l-head:] {
			return 2 + longestDecomposition(text[head:l-head])
		}
	}
	// 最后中间剩1个那就加个1，或者是剩1个大单词
	return 1
}

func main() {
	fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi")) //7
	fmt.Println(longestDecomposition("mercant"))                          //1
	fmt.Println(longestDecomposition("antaprezatepzapreanta"))            //11
	fmt.Println(longestDecomposition("antta"))                            //3
	fmt.Println(longestDecomposition("attta"))                            //5
	fmt.Println(longestDecomposition("a"))                                //1
	fmt.Println(longestDecomposition("elvtoelvto"))                       //2
	fmt.Println(longestDecomposition("aaaa"))                             //4
	fmt.Println(longestDecomposition("aaa"))                              //3
	fmt.Println(longestDecomposition("aa"))                               //2
}
