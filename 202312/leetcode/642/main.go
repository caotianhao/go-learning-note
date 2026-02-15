package main

import "fmt"

func diStringMatch(s string) []int {
	l, slice642, ii := len(s), make([]int, 0), 0
	dd := l
	for i := 0; i < l; i++ {
		if s[i] == 'I' {
			slice642 = append(slice642, ii)
			ii++
		} else {
			slice642 = append(slice642, dd)
			dd--
		}
	}
	slice642 = append(slice642, dd)
	return slice642
}

func main() {
	fmt.Println(diStringMatch("IDIIIDDDIDIDIDI"))
}
