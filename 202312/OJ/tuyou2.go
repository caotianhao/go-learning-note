package main

import "fmt"

func vie(s string) (res []int) {
	l := len(s)
	for i := 0; i < l; i++ {
		hm := make(map[string]struct{})
		k := deleteByInd(s, i)
		right := len(k) / 3
		k = k[:3*right]
		for j := 0; j < len(k); j += 3 {
			hm[k[j:j+3]] = struct{}{}
		}
		res = append(res, len(hm))
	}
	return
}

func deleteByInd(s string, idx int) string {
	r := []rune(s)
	var res string
	for i := 0; i < len(r); i++ {
		if i == idx {
			continue
		}
		res += string(s[i])
	}
	return res
}

func main() {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	var s string
	_, _ = fmt.Scanf("%s", &s)
	res := vie(s)
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}
