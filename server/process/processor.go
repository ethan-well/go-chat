package process

import (
	"fmt"
	common "go-chat/common/message"
	"go-chat/server/model"
	"go-chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 处理消息
// 根据消息的类型，使用对应的处理方式
func (this *Processor) messageProcess(message common.Message) (err error) {
	switch message.Type {
	case common.LoginMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserLogin(message.Data)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
		}
	case common.RegisterMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserRegister(message.Data)
		if err != nil {
			fmt.Printf("some error when register: %v\n", err)
		}
	case common.UserSendGroupMessageType:
		fmt.Println("user send group message!")
		gmp := GroupMessageProcess{}
		gmp.sendToGroupUsers(message.Data)
	case common.ShowAllOnlineUsersType:
		olP := OnlineInfoProcess{this.Conn}
		err = olP.showAllOnlineUserList()
		if err != nil {
			fmt.Printf("get all online user list error: %v\n", err)
		}
	case common.PointToPointMessageType:
		fmt.Println("point to point comminite!")
		pop := PointToPointMessageProcess{}
		err = pop.sendMessageToTargetUser(message.Data)
		var code int
		if err != nil {
			code = 400
		} else {
			code = 100
		}

		// responseClient(conn net.Conn, code int, data string, err error) {
		err := pop.responseClient(this.Conn, code, "", err)
		if err != nil {
			fmt.Printf("some err when popmessage: %v", err)
		}

	default:
		fmt.Printf("other type\n")
	}
	return
}

// 处理和用户的之间的通讯
func (this *Processor) MainProcess() {

	// 循环读来自客户端的消息
	for {
		dispatcher := utils.Dispatcher{Conn: this.Conn}
		message, err := dispatcher.ReadData()
		if err != nil {
			if err == io.EOF {
				cc := model.ClientConn{}
				cc.Del(this.Conn)
				fmt.Printf("client closed!\n")
				break
			}
			fmt.Printf("get login message error: %v", err)
		}

		// 处理来客户端的消息
		// 按照消息的类型，使用不同的处理方法
		err = this.messageProcess(message)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
			break
		}
	}
}
