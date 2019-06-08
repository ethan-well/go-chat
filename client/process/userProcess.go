package process

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
)

type UserProcess struct{}

func (up UserProcess) Login(userName, password string) (err error) {
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
		return
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
	dispatcher := utils.Dispatcher{Conn: conn}

	responseMsg, err = dispatcher.ReadDate()
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

// 处理用户注册
func (up UserProcess) Register(userName, password, password_confirm string) (err error) {
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
	dispatcher := utils.Dispatcher{Conn: conn}
	responseMsg, err = dispatcher.ReadDate()
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
