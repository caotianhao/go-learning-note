package main

import (
	"bytes"
	"fmt"
	"time"
)

const add = 1 << 18

// 对于简单而少量的拼接，使用运算符+和+=的效果虽然很好，但随着拼接操作次数的增加，
// 这种做法的效率并不高。如果需要在循环中拼接字符串，则使用空的字节缓冲区来拼接效率更高。
func main() {
	origin := "o"
	b := time.Now()
	for i := 0; i < add; i++ {
		origin += "a"
	}
	e := time.Since(b)
	fmt.Println("use +=", e)
	//fmt.Println(origin)

	buffer := bytes.Buffer{}
	buffer.WriteString("o")
	b = time.Now()
	for i := 0; i < add; i++ {
		buffer.WriteString("a")
	}
	e = time.Since(b)
	fmt.Println("buffer", e)
	//fmt.Println(buffer.String())
}
