package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过 go 向 redis 写入和读取数据
	//1.先链接到 redis
	//在之前需要启动 redis-server.exe
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn dial err =", err)
		return
	}
	//记得使用 defer 关闭
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close() err =", err)
			return
		}
	}(conn)
	//fmt.Println("conn success", conn)

	//2.写入数据
	_, err = conn.Do("set", "name", "caotianhao曹")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	//fmt.Println("set success")

	//3.读取数据
	//这样返回的 str1 是接口类型，不能通过类型断言转换
	//str1, err := conn.Do("get", "name")
	//应该用 redis.String
	str1, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	fmt.Println("get success", str1)

	//4.hash
	//_, err = conn.Do("hSet", "user001", "name", "cth")
	//if err != nil {
	//	fmt.Println("hSet err =", err)
	//	return
	//}
	//_, err = conn.Do("hSet", "user001", "age", 18)
	//if err != nil {
	//	fmt.Println("hSet err =", err)
	//	return
	//}
	//_, err = conn.Do("hSet", "user001", "salary", 12345.67)
	//if err != nil {
	//	fmt.Println("hSet err =", err)
	//	return
	//}
	//user001Name, err := redis.String(conn.Do("hGet", "user001", "name"))
	//if err != nil {
	//	fmt.Println("hGet err =", err)
	//	return
	//}
	//fmt.Println("hGet success", user001Name)
	//user001Age, err := redis.Int(conn.Do("hGet", "user001", "age"))
	//if err != nil {
	//	fmt.Println("hGet err =", err)
	//	return
	//}
	//fmt.Println("hGet success", user001Age)
	//user001Salary, err := redis.Float64(conn.Do("hGet", "user001", "salary"))
	//if err != nil {
	//	fmt.Println("hGet err =", err)
	//	return
	//}
	//fmt.Println("hGet success", user001Salary)
	//
	////5.批量操作 redis.Strings
	//_, err = conn.Do("mSet", "go_name", "golang", "go_address", "cth", "go_version", 1.17, "go_year", 2022)
	//if err != nil {
	//	fmt.Println("mSet err =", err)
	//	return
	//}
	//goStr, err := redis.Strings(conn.Do("mGet", "go_name", "go_address", "go_year", "go_version"))
	//if err != nil {
	//	fmt.Println("mGet err =", err)
	//	return
	//}
	//for _, v := range goStr {
	//	fmt.Println("mGet some success", v)
	//}
}
