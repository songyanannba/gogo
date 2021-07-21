package model

import (
	"net"
	"room/commom/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
