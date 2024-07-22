package main

import (
	"fmt"
	"reflect"
)

type t1 struct {
	name string
	age  int
}

type t2 struct {
	name string
	age  int
}

func main() {
	t := t1{"c", 25}
	fmt.Printf("%v\n", t)  //{c 25}
	fmt.Printf("%+v\n", t) //{name:c age:25}
	tt := t2{"b", 111}
	fmt.Println(reflect.TypeOf(t), reflect.TypeOf(t).Kind())   //main.t1 struct
	fmt.Println(reflect.TypeOf(tt), reflect.TypeOf(tt).Kind()) //main.t2 struct
	str := "사랑해"
	fmt.Println(len(str)) // 3*3=9
}
