package controller

import (
	"gogogo/202408/chatroom/model"
	"gogogo/202408/chatroom/utils"
	"log"
	"strconv"
)

type Uu model.User

func (u *Uu) String() string {
	return u.UserIP + ", ID:" + strconv.Itoa(u.UserID) + ", 加入于:" + u.JoinTime.Format("2006-01-02 15:04:05")
}

func Broadcast() {
	for {
		select {
		case user := <-utils.OnlineChan:
			utils.UserList[user] = struct{}{}
		case user := <-utils.OfflineChan:
			delete(utils.UserList, user)
			close(user.MessageQueue)
		case msg := <-utils.GlobalMsgQueue:
			if msg.TalkTo == "" {
				for user := range utils.UserList {
					if user.UserID == msg.Talker {
						continue
					}
					user.MessageQueue <- msg.Content
				}
			} else {
				flag := false
				for user := range utils.UserList {
					id := strconv.Itoa(user.UserID)
					if id == msg.TalkTo {
						user.MessageQueue <- msg.Content
						flag = true
					}
				}
				if !flag {
					log.Println("用户", msg.TalkTo, "不存在!")
				}
			}
		}
	}
}
