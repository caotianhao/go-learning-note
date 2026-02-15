package main

import (
	"database/sql"
	"fmt"
	//这个没用这个包，但是也得导入并且用杠划掉
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//用户名:密码@tcp(ip:端口)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/leetcode"
	db, err := sql.Open("mysql", dsn)
	//这里只检查 dsn 是否符合格式，当它解析不出对应的字段就会报这个错
	if err != nil {
		fmt.Println("格式有误", err)
		return
	}
	//db.Ping() 才会检查用户名密码正不正确
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	fmt.Println("连接数据库成功。")
}
