package main

import (
	"gogogo/202408/chatroom/controller"
	"gogogo/202408/chatroom/handler"
	"gogogo/202408/chatroom/utils"
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
