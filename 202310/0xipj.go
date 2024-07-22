package main

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	var (
		i        = 1
		i2 int32 = 2
		i3 int64 = 3
	)
	fmt.Println(reflect.TypeOf(i), unsafe.Sizeof(i))   // int 8
	fmt.Println(reflect.TypeOf(i2), unsafe.Sizeof(i2)) // int32 4
	fmt.Println(reflect.TypeOf(i3), unsafe.Sizeof(i3)) // int64 8
	// 0xipj 表示 i*2^j，类型为 float64
	i4 := 0x3p3
	i5 := http.StatusOK - 176
	i6 := "12"
	fmt.Println(int(math.Ceil(i4)) == i5)         // true
	fmt.Println(strconv.ParseInt(i6, 1<<4, 1<<6)) // 18 nil
}
