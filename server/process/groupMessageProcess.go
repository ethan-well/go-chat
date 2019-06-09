package process

import (
	"fmt"
	"net"
)

type GroupMessageProcess struct {
	Conn    net.Conn
	GroupID int
}

// 向组内不人员发送消息
func (gmp GroupMessageProcess) sendToGroupUsers(message string) {
	fmt.Printf("send group message process!")
}
