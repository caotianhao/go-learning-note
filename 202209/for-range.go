package main

import "fmt"

func main() {
	str := "hello world 韩顺平"
	for index, val := range str {
		fmt.Printf("index = %d, val = %c\n", index, val)
	}
}
