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
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
}
