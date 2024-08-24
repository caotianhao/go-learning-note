package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 查找最长子字符串
 * @param input string字符串 待查找的字符串
 * @return string字符串一维数组
 */
func matchMaxLengthSubString(input string) []string {
	hashMap := map[byte]int{}
	l := len(input)
	res := make([]string, 0)
	right, maxLen := -1, 0
	for i := 0; i < l; i++ {
		if i != 0 {
			delete(hashMap, input[i-1])
		}
		for right+1 < l && hashMap[input[right+1]] == 0 {
			hashMap[input[right+1]]++
			right++
		}
		maxLen = max(maxLen, right-i+1)
	}
	hashMap2 := map[byte]int{}
	right2 := -1
	for i := 0; i < l; i++ {
		if i != 0 {
			delete(hashMap2, input[i-1])
		}
		for right2+1 < l && hashMap2[input[right2+1]] == 0 {
			hashMap2[input[right2+1]]++
			right2++
		}
		if right2-i+1 == maxLen {
			res = append(res, input[i:right2+1])
		}
	}
	return res
}

func main() {
	fmt.Println(matchMaxLengthSubString("sheinerintheworld"))
}
