package process

import (
	"fmt"
	commen "go-chat/commen/message"
)

// 处理消息
// 根据消息的类型，使用对应的处理方式
func MessgeProcess(message commen.Message) (code int, err error) {
	switch message.Type {
	case commen.LoginMessageType:
		code, err = userLogin(message.Data)
	case commen.ResponseMessageType:
		fmt.Println(commen.ResponseMessageType)
	default:
		fmt.Printf("other type")
	}
	return
}
