package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var str1 = "true"
	var b bool
	//有两个返回值
	b, err = strconv.ParseBool(str1)
	if err != nil {
		return
	}
	fmt.Printf("b type is %T, b = %v\n", b, b)

	var str2 = "12345"
	var i int64
	i, err = strconv.ParseInt(str2, 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("i type is %T, i = %v\n", i, i)

	var str3 = "123.4567"
	var f float64
	f, err = strconv.ParseFloat(str3, 64)
	if err != nil {
		return
	}
	fmt.Printf("f type is %T, f = %v\n", f, f)
	//fmt.Println(err)

	//string转基本类型时，要注意可以转才行
	//如果是这种，不会报错，但也不会转成97，会是0
	var str4 = "a"
	var j int64
	j, err = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("j type is %T, j = %v\n", j, j)
}
