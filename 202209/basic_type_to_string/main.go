package main

import (
	"fmt"
	"strconv"
	"time"
)

func getB() bool {
	return false
}

func main() {
	var i int64 = 64
	var f = 1.123
	var b = getB()
	var c byte = 'a'
	var str string

	//第一种方式，用fmt.Sprintf
	str = fmt.Sprintf("%d", i)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", f)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	//%t is bool
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	fmt.Println("Now is", time.Now())

	//第二种方式，用strconv包
	//10表示十进制
	str = strconv.FormatInt(i, 10)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	//'f'表示格式
	//10表示保留10位小数
	//64表示float64
	str = strconv.FormatFloat(f, 'f', 10, 64)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type is %T, str = %q\n", str, str)

	//strconv包中还有一个函数可以使int转为string
	var j = 888
	str = strconv.Itoa(j)
	fmt.Printf("str type is %T, str = %q\n", str, str)
}
