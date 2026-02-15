package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	bool    // 1
	float64 // 8
	int16   // 2
}

type s2 struct {
	float64
	int16
	bool
}

type s3 struct {
	bool
	int16
	float64
}

type s4 struct {
	int16
	float64
	bool
}

type s5 struct {
	int16
	float64
	int
	bool
}

func main() {
	var (
		ss1 s1
		ss2 s2
		ss3 s3
		ss4 s4
		ss5 s5
	)
	fmt.Println("bool =", unsafe.Sizeof(true))
	fmt.Println("float64 =", unsafe.Sizeof(1.23))
	fmt.Println("int16 =", unsafe.Sizeof(int16(1)))
	fmt.Println("ss1 =", unsafe.Sizeof(ss1), unsafe.Alignof(ss1))
	fmt.Println("ss2 =", unsafe.Sizeof(ss2), unsafe.Alignof(ss2))
	fmt.Println("ss3 =", unsafe.Sizeof(ss3), unsafe.Alignof(ss3))
	fmt.Println("ss4 =", unsafe.Sizeof(ss4), unsafe.Alignof(ss4))
	fmt.Println("ss5 =", unsafe.Sizeof(ss5), unsafe.Alignof(ss5))
}
