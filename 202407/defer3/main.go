package main

import "fmt"

func defer3() {
	fmt.Println("defer3() ++++++")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
			fmt.Println("defer3() ------")
		}
	}()
	panic("test panic")
}

func main() {
	defer3()
}
