package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义全局连接池
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}, //初始化连接代码
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //最大连接数，0 表示无限制
		IdleTimeout: 100, //在此持续时间内保持空闲后关闭连接
	}
}

func main() {
	//先从 redis 连接池中取出一个连接
	//如果要取出连接，一定要保证它没关闭
	conn := pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close() err =", err)
			return
		}
	}(conn)

	//继续操作
	_, err := conn.Do("set", "clock", "2022.10.25")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	str, err := redis.String(conn.Do("get", "clock"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	fmt.Println("get success", str)

	//如果要取出连接，一定要保证它没关闭
	err = pool.Close()
	if err != nil {
		fmt.Println("pool.Close() err =", err)
		return
	}
	conn2 := pool.Get()
	_, err = conn2.Do("set", "clock2", "2022.10.25.25")
	if err != nil {
		//报错：redigo: get on closed pool
		fmt.Println("set err =", err)
		return
	}
	fmt.Println("pool.Close(), conn2 := pool.Get() success")
}
