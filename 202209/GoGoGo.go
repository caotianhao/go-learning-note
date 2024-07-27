package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	c := [3]int{5, 6}
	if a == b {
		fmt.Println("a == b")
	}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
	fmt.Println(a, b, c)

	d := [...]int{5, 6, 7}
	fmt.Println(len(d), cap(d), d)
}
