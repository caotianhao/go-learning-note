package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	strNum := strconv.FormatInt(int64(num), 10)
	l, cnt := len(strNum), 0
	for i := 0; i < l-k; i++ {
		temp, _ := strconv.ParseInt(strNum[i:i+k], 10, 64)
		tmp := int(temp)
		if tmp == 0 {
			cnt += 0
		} else if num%tmp == 0 {
			cnt++
		}
	}
	t1, _ := strconv.ParseInt(strNum[l-k:], 10, 64)
	t2 := int(t1)
	if t2 == 0 {
		cnt += 0
	} else if num%t2 == 0 {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(divisorSubstrings(240, 2))    //2
	fmt.Println(divisorSubstrings(430043, 2)) //2
	fmt.Println(divisorSubstrings(10001, 4))  //2
}
