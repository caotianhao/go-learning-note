package main

import (
	"chatroom/utils"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", utils.LocalAddr)

	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("已结束!")
		done <- struct{}{}
	}()

	_, _ = io.Copy(conn, os.Stdin)
	_ = conn.Close()
	<-done
}
