package main

import (
	"fmt"
	"strconv"
)

// 9
// a2b3cde3f
// abbcccdefff
func deRar(s string) (res string) {
	lastIdx := 0
	stringList := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if isChar(s[i]) {
			stringList = append(stringList, s[lastIdx:i+1])
			lastIdx = i + 1
		}
	}
	for _, v := range stringList {
		l := len(v)
		if l == 1 {
			res += string(v[0])
		} else {
			cnt := v[:l-1]
			num, _ := strconv.Atoi(cnt)
			for i := 0; i < num; i++ {
				res += string(v[l-1])
			}
		}
	}
	return
}

func isChar(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func main() {
	lengthOfString := 0
	_, _ = fmt.Scanf("%d", &lengthOfString)
	var str string
	_, _ = fmt.Scanf("%s", &str)
	fmt.Println(deRar(str))
}
