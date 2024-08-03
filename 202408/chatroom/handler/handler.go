package handler

import (
	"bufio"
	"chatroom/controller"
	"chatroom/model"
	"chatroom/utils"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func HandleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	user := &controller.Uu{
		UserID:       getID(),
		UserIP:       conn.RemoteAddr().String(),
		JoinTime:     time.Now(),
		MessageQueue: make(chan string, 10),
	}

	go sendMessage(conn, user.MessageQueue)

	user.MessageQueue <- "欢迎 " + user.String()

	msg := &model.Message{
		Talker:  user.UserID,
		Content: "用户" + strconv.Itoa(user.UserID) + "加入了",
	}

	utils.GlobalMsgQueue <- msg
	utils.OnlineChan <- (*model.User)(user)

	active := make(chan struct{})
	timer := time.NewTimer(utils.Duration)

	go func() {
		for {
			select {
			case <-timer.C:
				_ = conn.Close()
			case <-active:
				timer.Reset(utils.Duration)
			}
		}
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		line := input.Text()
		fmt.Println(line)
		if strings.HasPrefix(line, "@") {
			msg.TalkTo = strings.SplitN(line, " ", 2)[0][1:]
			line = strings.SplitN(line, " ", 2)[1]
		} else {
			msg.TalkTo = ""
		}
		msg.Content = "User-" + strconv.Itoa(user.UserID) + ": " + line
		utils.GlobalMsgQueue <- msg
	}
	if err := input.Err(); err != nil {
		log.Fatalln("读取失败！", err)
	}

	utils.OfflineChan <- (*model.User)(user)

	msg.Content = "用户" + strconv.Itoa(user.UserID) + "离开了"
	utils.GlobalMsgQueue <- msg
}

func sendMessage(conn net.Conn, msgQueue <-chan string) {
	for msg := range msgQueue {
		_, _ = fmt.Fprintln(conn, msg)
	}
}

func getID() int {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	utils.Uid++
	return utils.Uid
}
