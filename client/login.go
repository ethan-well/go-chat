package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	commen "go-chat/commen/message"
	"net"
)

func readDate(conn net.Conn) (msg commen.ResponseMessage, err error) {
	buf := make([]byte, 10240)

	// 读取消息长度信息
	n, err := conn.Read(buf[:4])
	if err != nil {
		return
	}
	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(buf[0:4])

	// 读取消息本身
	n, err = conn.Read(buf[:dataLen])
	if err != nil {
		fmt.Printf("server read data login data error: %v", err)
	}

	// 对比消息本身的长度和期望长度是否匹配
	if n != int(dataLen) {
		err = errors.New("login message length error")
		return
	}

	// 从 conn 中解析消息并存放到 msg 中，此处一定传递的是 msg 的地址
	err = json.Unmarshal(buf[:dataLen], &msg)
	if err != nil {
		fmt.Printf("json.Unmarshl error: %v", err)
	}
	return
}

func login(userName, password string) (err error) {
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
	loginMessage.UserName = userName
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
		return
	}

	//客户端发送消息本身
	writeLen, err = conn.Write(data)
	if err != nil {
		fmt.Printf("send data length to server error: %v", err)
		return
	}

	// 接受服务端返回
	var responseMsg commen.ResponseMessage
	responseMsg, err = readDate(conn)
	if err != nil {
		fmt.Printf("some error, retry please!\n")
		return
	}

	switch responseMsg.Code {
	case 200:
		fmt.Printf("Loggin succeed!")
	case 500:
		err = errors.New("server error")
	case 404:
		err = errors.New("user not exist")
	case 403:
		err = errors.New("pasword not valide")
	default:
		err = errors.New("some error")
	}

	return
}
