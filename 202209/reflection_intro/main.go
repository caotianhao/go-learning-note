package main

import (
	"fmt"
	"reflect"
)

func reflectTest(a any) {

	//通过反射获取到 num 的 type
	rType := reflect.TypeOf(a)
	fmt.Printf("rType is %T, ", rType)
	fmt.Printf("rType = %v\n", rType)

	//通过反射获取到 num 的值
	rValue := reflect.ValueOf(a)
	fmt.Printf("rValue is %T, ", rValue)
	fmt.Printf("rValue = %v\n", rValue)

	//int64 型变量
	var n int64 = 9
	//不能通过相加得到，也就是不能用 int64 加 reflect.Value 得到结果
	//fmt.Println(n + rValue)
	//要使用 rValue.Int() 转化为 int， 但注意返回值是 int64 类型
	fmt.Println(n + rValue.Int())

	//先将 rValue 转成接口
	trueValue := rValue.Interface()
	//然后再通过类型断言转为原值
	fmt.Printf("trueValue is %T, trueValue = %v\n", trueValue, trueValue)
	fmt.Println(trueValue.(int))
}

func main() {
	num := 10
	reflectTest(num)
	//fmt.Printf("normal num is %T\n", num)
}
