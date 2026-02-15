package main

import (
	"fmt"
	"strings"
)

// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

/*
思路：如果字符串 s 包含一个重复的子字符串，那么这意味着可以多次移位和换行该字符串，并使其与原始字符串匹配。
例如：abcabc
移位一次：cabcab
移位两次：bcabca
移位三次：abcabc
现在字符串和原字符串匹配了，所以可以得出结论存在重复的子串。
基于这个思想，可以每次移动 k 个字符，直到匹配移动 length-1 次。
但是这样对于重复字符串很长的字符串，效率会非常低，执行时间超时了。
为了避免这种无用的环绕，可以创建一个新的字符串 str，它等于原来的字符串 s 再加上 s 自身，这样其实就包含了所有移动的字符串。
比如字符串：S = acd，那么 str = S + S = acdacd
acd 移动的可能：dac、cda。其实都包含在了 str 中了，就像一个滑动窗口
一开始 acd (acd) ，移动一次 ac(dac)d，移动两次 a(cda)cd。循环结束
所以可以直接判断 str 中去除首尾元素之后，是否包含自身元素，如果包含，则表明存在重复子串。
*/
func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s+s)-1], s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
