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

func (cc ClientConn) Del(userConn net.Conn) {
	for id, conn := range ClientConnsMap {
		if conn == userConn {
			delete(ClientConnsMap, id)
		}
	}
}
