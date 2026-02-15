package main

import (
	"fmt"
	"time"
)

func foo1(n uint) int {
	if n <= 1 {
		return 1
	}
	s := make([]int, n)
	s[0], s[1] = 1, 1
	for i := uint(2); i < n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n-1]
}

func foo2(n uint) int {
	if n <= 1 {
		return 1
	}
	s := [2]int{1, 1}
	z := 0
	for i := uint(2); i < n; i++ {
		s[z] = s[0] + s[1]
		z ^= 1
	}
	return s[z^1]
}

func foo3(n uint) int {
	if n <= 1 {
		return 1
	}
	s := [1 << 25]int{1, 1}
	for i := uint(2); i < n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n-1]
}

func main() {
	var s uint = 9999999
	t := time.Now()
	foo1(s)
	t1 := time.Now()
	tt := t1.Sub(t)
	fmt.Println("foo1", tt)

	t = time.Now()
	foo2(s)
	t1 = time.Now()
	tt = t1.Sub(t)
	fmt.Println("foo2", tt)

	t = time.Now()
	foo3(s)
	t1 = time.Now()
	tt = t1.Sub(t)
	fmt.Println("foo3", tt)
}
