package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p unsafe.Pointer
	var i, j int
	fmt.Println(unsafe.Sizeof(p), unsafe.Sizeof(i), unsafe.Sizeof(j))
	var s string
	fmt.Println("string", unsafe.Sizeof(s))
	var p2 uint8
	fmt.Println("uintptr", unsafe.Sizeof(p2))
}
