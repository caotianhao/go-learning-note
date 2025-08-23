package main

import (
	"fmt"
	"math/bits"
)

func countPoints(rings string) int {
	ring, l, cnt := make([]string, 10), len(rings), 0
	for i := 0; i < l; i += 2 {
		ring[rings[i+1]-48] += string(rings[i])
	}
	flagR, flagG, flagB := false, false, false
	//fmt.Println(ring)
	for i := 0; i < 10; i++ {
		flagR, flagG, flagB = false, false, false
		if isCharInString(ring[i], 'R') {
			flagR = true
		}
		if isCharInString(ring[i], 'G') {
			flagG = true
		}
		if isCharInString(ring[i], 'B') {
			flagB = true
		}
		if flagR && flagG && flagB {
			cnt++
		}
		//if i == 1 || i == 2 {
		//	fmt.Println(flagR, flagG, flagB)
		//}
	}
	return cnt
}

func isCharInString(s string, c byte) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		if c == s[i] {
			return true
		}
	}
	return false
}

func countPoints1(rings string) int {
	r, g, b, n := 0, 0, 0, len(rings)
	for i := 0; i < n; i += 2 {
		if rings[i] == 'R' {
			r |= 1 << int(rings[i+1]-48)
		} else if rings[i] == 'G' {
			g |= 1 << int(rings[i+1]-48)
		} else {
			b |= 1 << int(rings[i+1]-48)
		}
	}
	return bits.OnesCount64(uint64(r & g & b))
}

func main() {
	//fmt.Println(countPoints("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints1("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints("B0B6G0R6R0R6G9G8B5B9B7R5G2B3B9R4G5B1B5G0B5"))
	fmt.Println(countPoints1("B0B6G0R6R0R6G9G8B5B9B7R5G2B3B9R4G5B1B5G0B5"))
	fmt.Println(countPoints("B0B6G0R6R0R6G9G8B5B9B7R5G2B3B9R4G5B1B5G0B5G5R5B6G6R6B7B8B9R7R8R9G5G6G7G8G9"))
	fmt.Println(countPoints1("B0B6G0R6R0R6G9G8B5B9B7R5G2B3B9R4G5B1B5G0B5G5R5B6G6R6B7B8B9R7R8R9G5G6G7G8G9"))
	//fmt.Println(isCharInString("bg", 'a'))
}
