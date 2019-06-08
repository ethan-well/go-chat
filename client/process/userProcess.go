package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
	"os"
)

type UserProcess struct{}

// 登陆成功菜单显示：
func showAfterLoginMenu() {

	fmt.Println("----------------login succeed!----------------")
	fmt.Println("\t\tselect what you want to do")
	fmt.Println("\t\t1. Show all online users")
	fmt.Println("\t\t2. Send message")
	fmt.Println("\t\t3. Show history messsge")
	fmt.Println("\t\t4. exist")
	var key int
	var content string

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("Show all online users")
	case 2:
		fmt.Println("Say some thing!")
		fmt.Scanf("%s\n", &content)
	case 3:
		fmt.Println("Message List：")
	case 4:
		fmt.Println("Exit...")
		os.Exit(0)
	default:
		fmt.Println("selected invalied!")
	}
}

// 用户登陆
func (up UserProcess) Login(userName, password string) (err error) {
	// 链接服务器
	conn, err := net.Dial("tcp", "localhost:8888")

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

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)
	if err != nil {
		return
	}

	// 接受服务端返回
	c := make(chan bool, 1)
	resErr := make(chan error, 1)
	go Response(conn, c, resErr)
	<-c
	if err = <-resErr; err != nil {
		return
	}

	for {
		showAfterLoginMenu()
	}
}

// 处理用户注册
func (up UserProcess) Register(userName, password, password_confirm string) (err error) {
	if password != password_confirm {
		err = errors.New("confirm password not match")
		return
	}

	conn, err := net.Dial("tcp", "localhost:8888")

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

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)
	if err != nil {
		fmt.Printf("send data erro!\n")
		return
	}

	// 接受服务端返回
	// 避免主 routine 提前返回
	c := make(chan bool, 1)

	resErr := make(chan error, 1)
	go Response(conn, c, resErr)
	<-c
	err = <-resErr

	return
}
