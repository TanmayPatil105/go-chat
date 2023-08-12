package router

import (
	"time"
)

type Message struct {
	UserId    string
	RoomId    string
	Text      string
	TimeStamp time.Time
}
