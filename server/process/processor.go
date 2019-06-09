package process

import (
	"fmt"
	commen "go-chat/commen/message"
	"go-chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 处理消息
// 根据消息的类型，使用对应的处理方式
func (this *Processor) messgeProcess(message commen.Message) (err error) {
	switch message.Type {
	case commen.LoginMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserLogin(message.Data)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
		}
	case commen.RegisterMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserRegister(message.Data)
		if err != nil {
			fmt.Printf("some error when register: %v\n", err)
		}
	case commen.UserSendGroupMessageType:
		fmt.Printf("user send group message!")
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
				fmt.Printf("client closed!\n")
				break
			}
			fmt.Printf("get login message error: %v", err)
		}

		// 处理来客户端的消息
		// 按照消息的类型，使用不同的处理方法
		err = this.messgeProcess(message)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
			break
		}
	}
}
