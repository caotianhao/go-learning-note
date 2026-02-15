package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) (res []string) {
	cntMap := map[string]int{}
	for _, v := range cpdomains {
		countDot := strings.Count(v, ".")
		if countDot == 2 {
			space := strings.Index(v, " ")
			left := strings.IndexFunc(v, func(r rune) bool {
				return r == '.'
			})
			right := strings.LastIndex(v, ".")
			cnt, _ := strconv.ParseInt(v[:space], 10, 64)
			cntMap[v[space+1:]] += int(cnt)
			cntMap[v[left+1:]] += int(cnt)
			cntMap[v[right+1:]] += int(cnt)
		} else {
			space := strings.Index(v, " ")
			dot := strings.LastIndex(v, ".")
			cnt, _ := strconv.ParseInt(v[:space], 10, 64)
			cntMap[v[space+1:]] += int(cnt)
			cntMap[v[dot+1:]] += int(cnt)
		}
	}
	for i, v := range cntMap {
		res = append(res, strconv.FormatInt(int64(v), 10)+" "+i)
	}
	return
}

func main() {
	arr := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(arr))
}
