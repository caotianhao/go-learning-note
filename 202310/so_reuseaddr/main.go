package main

import (
	"net"
)

func main() {
	// SO_REUSEADDR（Socket Option Reuse Address）
	// 是一种套接字选项，通常在编程中用于控制套接字的行为
	// 主要作用是允许多个套接字绑定到同一端口，即使该端口已被一个套接字占用
	// 主要应用场景是在网络编程中，特别是在服务器编程中，
	// 它允许服务器在停止后快速重新启动并绑定到相同的端口，而不需要等待操作系统释放该端口
	// 这可以提高服务器的可用性和容错性。
	// net.Listen 默认开启 SO_REUSEADDR
	listener, _ := net.Listen("tcp", ":8080")
	_ = listener.Close()
}
