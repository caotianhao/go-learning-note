package main

import (
	"fmt"
	"strings"
)

func kunHao(str string) string {
	return strings.Replace(strings.Replace(str, "]", "}", -1), "[", "{", -1)
}

func yinHao(str string) string {
	return strings.Replace(str, "\"", "'", -1)
}

func main() {
	fmt.Println(kunHao("[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]"))
	fmt.Println(yinHao("\"''a"))
}
