package model

import (
	"net"
)

type ClientConn struct{}

var ClientConnsMap map[int]net.Conn

func init() {
	ClientConnsMap = make(map[int]net.Conn)
}

func (cc ClientConn) Save(userID int, userConn net.Conn) {
	ClientConnsMap[userID] = userConn
}
