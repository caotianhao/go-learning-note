package main

import "fmt"

func moveSubStr(str string, n int) string {
	return str[n-1:] + str[:n-1]
}

func main() {
	fmt.Println(moveSubStr("abcdefg", 3))
}
