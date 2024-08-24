package main

import "fmt"

func changeST(s, t string) string {
	//hm := make(map[byte]byte)
	hm := make(map[byte][]byte)
	hm2 := make(map[byte]struct{})
	n := len(s)
	nt := len(t)
	if n != nt {
		return "No"
	}
	for i := 0; i < n; i++ {
		//sv, ok := hm[s[i]]
		//if !ok {
		//	hm[s[i]] = t[i]
		//} else {
		//	if sv != t[i] {
		//		return "No"
		//	}
		//}
		if _, ok := hm2[t[i]]; !ok {
			hm[s[i]] = append(hm[s[i]], t[i])
			hm2[t[i]] = struct{}{}
		}

	}
	for _, v := range hm {
		if len(v) != 1 {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	var s, t string
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &s)
		_, _ = fmt.Scanf("%s", &t)
		fmt.Println(changeST(s, t))
	}
}
