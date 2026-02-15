package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 54
	fmt.Printf("sizeof a is %d, type of a is %T\n", unsafe.Sizeof(a), a)
	b := 123456789011
	fmt.Printf("sizeof b is %d, type of b is %T\n", unsafe.Sizeof(b), b)
	c := 1234567890113333366
	fmt.Printf("sizeof c is %d, type of c is %T\n", unsafe.Sizeof(c), c)
	var d int32 = 1
	fmt.Printf("sizeof d is %d, type of d is %T\n", unsafe.Sizeof(d), d)
	var e int64 = 1
	fmt.Printf("sizeof e is %d, type of e is %T\n", unsafe.Sizeof(e), e)
	var f int8 = 1
	fmt.Printf("sizeof f is %d, type of f is %T\n", unsafe.Sizeof(f), f)
	//rune = int32
	var g rune = 1
	fmt.Printf("sizeof g is %d, type of g is %T\n", unsafe.Sizeof(g), g)
}
