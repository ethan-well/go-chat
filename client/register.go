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
		fmt.Printf("client soem error: %v", err)
	}

	// 构造需要传递给服务器的数据
	messsage.Data = string(data)
	messsage.Type = commen.RegisterMessageType

	data, err = json.Marshal(messsage)
	if err != nil {
		fmt.Printf("registerMessage json Marshal error: %v", err)
		return
	}

	// 发送消息长度
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)

	// 发送消息长度
	_, err = conn.Write(bytes[:])
	if err != nil {
		fmt.Printf("some error")
		return
	}

	//客户端发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("send data length to server error: %v", err)
		return
	}

	// 接收服务器返回
	var responseMsg commen.ResponseMessage
	responseMsg, err = readDate(conn)
	if err != nil {
		fmt.Printf("some error, retry please!\n")
		return
	}

	switch responseMsg.Code {
	case 200:
		fmt.Printf("Register succeed!\n")
	case 500:
		err = errors.New("server error")
	case 403:
		err = errors.New("user has already existed!")
	case 402:
		err = errors.New("pasword not match!")
	default:
		err = errors.New("some error")
	}

	return
}
