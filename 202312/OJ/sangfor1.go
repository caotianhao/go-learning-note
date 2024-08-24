package main

import (
	"fmt"
	"strings"
)

func isLegal(domain string) bool {
	if strings.Count(domain, ".") < 1 {
		return false
	}
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}
	for i := 0; i < len(domain)-1; i++ {
		if domain[i] == '.' && domain[i+1] == '.' {
			return false
		}
	}
	part := strings.Split(domain, ".")
	l := 0
	for _, v := range part {
		if len(v) == 0 {
			return false
		}
		if v[0] == '-' || v[len(v)-1] == '-' {
			return false
		}
		for i := 1; i < len(v)-1; i++ {
			if !helpDomain(rune(v[i])) {
				return false
			}
		}
		//for _, vv := range v {
		//	if !helpDomain(vv) {
		//		return false
		//	}
		//}
		l += len(v)
		if len(v) > 1<<6-1 {
			return false
		}
	}
	if l > 1<<8-1 {
		return false
	}
	return true
}

func helpDomain(s rune) bool {
	return s >= '0' && s <= '9' || s >= 'A' && s <= 'Z' || s >= 'a' && s <= 'z' || s == '-'
}

func main() {
	var domain string
	_, _ = fmt.Scanf("%s", &domain)
	fmt.Println(isLegal(domain))
}
