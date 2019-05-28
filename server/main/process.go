package main

import (
	"fmt"
	commen "go-chat/commen/message"
)

// 处理消息
// 根据消息的类型，使用对应的处理方式
func process(message commen.Message) (code int, err error) {
	switch message.Type {
	case commen.LoginMessageType:
		code, err = dealWithLoginMessage(message.Data)
	case commen.ResponseMessageType:
		fmt.Println(commen.ResponseMessageType)
	default:
		fmt.Printf("other type")
	}
	return
}
