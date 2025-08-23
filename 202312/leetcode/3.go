package main

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//	q, cnt, ptr, res, map3 := -1, 0, 0, 0, map[byte]int{}
//	for ptr < len(s) {
//		//检查对应 ptr 的字符是否在 map 中
//		j, ok := map3[s[ptr]]
//		if !ok {
//			//如果不在就加入，值为下标
//			map3[s[ptr]] = ptr
//			//长度加一
//			cnt++
//			//继续
//			ptr++
//		} else {
//			//如果在的话，记录一下长度，因为串要变化了
//			res = max3(res, cnt)
//			//删除之前对应字符出现位置之前直至上上次删除的所有记录
//			for i := j; i >= q+1; i-- {
//				delete(map3, s[i])
//			}
//			//上上次删除位置更新
//			q = j
//			//加入
//			map3[s[ptr]] = ptr
//			//新长度
//			cnt = ptr - j
//			//继续
//			ptr++
//		}
//	}
//	//再次更新长度
//	res = max3(res, cnt)
//	return res
//}

func lengthOfLongestSubstring(s string) (res int) {
	l := len(s)
	p, q, hashMap := 0, 0, map[byte]struct{}{}
	for q < l {
		_, ok := hashMap[s[q]]
		if !ok {
			hashMap[s[q]] = struct{}{}
			res = max(res, len(hashMap))
			q++
		} else {
			delete(hashMap, s[p])
			p++
		}
	}
	return
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //3
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))   //1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   //3
	fmt.Println(lengthOfLongestSubstring("dsfhavx"))  //7
	fmt.Println(lengthOfLongestSubstring("a"))        //1
	fmt.Println(lengthOfLongestSubstring("qwq"))      //2
	fmt.Println(lengthOfLongestSubstring("QAQ"))      //2
	fmt.Println(lengthOfLongestSubstring("ppspp"))    //2
	////                                    0123456789
	fmt.Println(lengthOfLongestSubstring("mfdhzrcmrowvmxleitnenur")) //11
	fmt.Println(lengthOfLongestSubstring("afvsdqufia"))              //8
	fmt.Println(lengthOfLongestSubstring("sdadshhdffffffffffffffvsdqufisduh" +
		"auhsughiuhenvojcnpdepfealkbdgbfklengkjfdslh")) //10
	fmt.Println(lengthOfLongestSubstring("abcbdbbbabcbb"))
}
