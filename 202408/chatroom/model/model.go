package model

import (
	"time"
)

type User struct {
	UserID       int
	UserIP       string
	JoinTime     time.Time
	MessageQueue chan string
}

type Message struct {
	Talker  int
	Content string
	TalkTo  string
}
