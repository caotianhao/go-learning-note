package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxPrefixIP(ip1, ip2 string) int {
	d1 := strings.Split(ip1, ".")
	d2 := strings.Split(ip2, ".")
	var s1, s2 string
	for _, v := range d1 {
		t, _ := strconv.ParseInt(v, 10, 64)
		tt := strconv.FormatInt(t, 2)
		var ttt string
		n := len(tt)
		for i := 0; i < 8-n; i++ {
			ttt += "0"
		}
		ttt += tt
		//fmt.Println("1", ttt)
		s1 += ttt
	}
	for _, v := range d2 {
		t, _ := strconv.ParseInt(v, 10, 64)
		tt := strconv.FormatInt(t, 2)
		var ttt string
		n := len(tt)
		for i := 0; i < 8-n; i++ {
			ttt += "0"
		}
		ttt += tt
		//fmt.Println("2", ttt)
		s2 += ttt
	}
	cnt := 0
	for i := 0; i < 32; i++ {
		if s1[i] == s2[i] {
			cnt++
		} else {
			break
		}
	}
	//fmt.Println(s1, s2)
	return cnt
}

func main() {
	var ip1, ip2 string
	_, _ = fmt.Scanf("%s %s", &ip1, &ip2)
	fmt.Println(maxPrefixIP(ip1, ip2))
}
