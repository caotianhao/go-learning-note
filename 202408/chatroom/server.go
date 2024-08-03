package main

import (
	"chatroom/controller"
	"chatroom/handler"
	"chatroom/utils"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", utils.LocalAddr)

	go controller.Broadcast()

	for {
		conn, _ := listener.Accept()
		go handler.HandleConn(conn)
	}
}
