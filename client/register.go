package main

import (
	"encoding/binary"
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

	// 定义消息
	var messsage commen.Message

	// 生成 registerMessage
	var registerMessage commen.RegisterMessage
	registerMessage.UserName = userName
	registerMessage.Password = password
	registerMessage.PasswordConfirm = password_confirm

	data, err := json.Marshal(registerMessage)
	if err != nil {
		fmt.Printf("client soem error: %T")
	}

	// 构造需要传递给服务器的数据
	messsage.Data = string(data)
	messsage.Type = commen.RegisterMessageType
	data, _ = json.Marshal(messsage)

	// 发送消息长度
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)

	// 发送数据
	_, err = conn.Write(bytes[:])
	if err != nil {
		fmt.Printf("some error")
		return
	}

	// 接收服务器返回

	return
}
