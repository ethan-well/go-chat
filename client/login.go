package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	commen "go-chat/commen/message"
	"net"
)

func login(userID int, password string) (err error) {
	// 链接服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect server error: %v", err)
		return
	}

	var message commen.Message
	message.Type = commen.LoginMessageType
	// 生成 loginMessage
	var loginMessage commen.LoginMessage
	loginMessage.UserID = userID
	loginMessage.Password = password

	// func Marshal(v interface{}) ([]byte, error)
	// 先序列话需要传到服务器的数据
	data, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Printf("some error when parse you data, error: %v\n", err)
	}

	// 首先发送数据 data 的长度到服务器端
	// 将一个字符串的长度转为一个表示长度的切片
	message.Data = string(data)
	message.Type = commen.LoginMessageType
	data, _ = json.Marshal(message)

	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)

	// 客户端发送消息长度
	writeLen, err := conn.Write(bytes[:])
	if writeLen != 4 || err != nil {
		fmt.Printf("send data to server error: %v", err)
		return err
	}

	//客户端发送消息本身
	writeLen, err = conn.Write(data)
	if err != nil {
		fmt.Printf("send data length to server error: %v", err)
		return err
	}

	// 接受服务端的返回
	return nil
}
