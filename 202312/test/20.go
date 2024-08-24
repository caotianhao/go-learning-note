package main

import (
	"fmt"
	"gogogo/202312/test/t1"
)

type Student = t1.Student

func main() {
	s := &Student{
		Age:  1,
		Name: "aa",
	}
	fmt.Println(s.Age)
}
