package main

import (
	"encoding/json"
	"errors"
	"fmt"
	commen "go-chat/commen/message"
	"net"
)

// 处理用户注册
func register(userName, password, password_confirm string) (err error) {
	if password != password_confirm {
		err = errors.New("confirm password not match")
		return
	}

	conn, err := net.Dial("tcp", "localhost:8888")
	defer conn.Close()

	if err != nil {
		fmt.Printf("connect server error: %v", err)
		return
	}

	// 定义消息类型
	var messsage commen.Message
	messsage.Type = commen.RegisterMessageType

	// 生成 registerMessage
	var registerMessage commen.RegisterMessage
	registerMessage.UserName = userName
	registerMessage.Password = password
	registerMessage.PasswordConfirm = password_confirm

	data, err := json.Marshal(registerMessage)
	if err != nil {
		fmt.Printf("client soem error: %T")
	}

	return
}
