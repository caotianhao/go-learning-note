package main

import (
	"fmt"
	"reflect"
)

type ii = int

type ii2 int

func type1(a, b ii) ii {
	return a + b
}

func type2(a, b ii2) ii2 {
	return a + b
}

func main() {
	t1, t2 := type1(1, 2), type2(1, 2)
	fmt.Println(t1, t2)                                        // 3 3
	fmt.Println(reflect.TypeOf(t1), reflect.TypeOf(t1).Kind()) // int int
	fmt.Println(reflect.TypeOf(t2), reflect.TypeOf(t2).Kind()) // main.ii2 int
}
