package main

import "fmt"

func compressString(param string) (ans string) {
	n := len(param)
	if n == 0 {
		return
	}
	tmp := make([]string, 0)
	left, right := 0, 1
	for right < n {
		if param[left] != param[right] {
			tmp = append(tmp, string(param[left]))
			tmp = append(tmp, fmt.Sprintf("%d", right-left))
			left = right
		}
		right++
	}
	tmp = append(tmp, string(param[left]), fmt.Sprintf("%d", right-left))
	for _, v := range tmp {
		if v != "1" {
			ans += v
		}
	}
	return
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
}
