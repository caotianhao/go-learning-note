package utils

import (
	"chatroom/model"
	"sync"
	"time"
)

const (
	LocalAddr = "localhost:9725"
	Capacity  = 1 << 6
	Duration  = time.Hour
)

var (
	UserList       = make(map[*model.User]struct{})
	OnlineChan     = make(chan *model.User)
	OfflineChan    = make(chan *model.User)
	GlobalMsgQueue = make(chan *model.Message, Capacity)
	Uid            int
	Mu             sync.Mutex
)
