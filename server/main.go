package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	commen "go-chat/commen/message"
	"io"
	"net"
)

func readDate(conn net.Conn) (msg commen.Message, err error) {
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

func login(userID int, passWord string) bool {
	// 判断用户名和密码
	return userID == 100 && passWord == "123"
}

func dealWithLoginMessage(message string) {
	var info commen.LoginMessage
	err := json.Unmarshal([]byte(message), &info)
	if err != nil {
		fmt.Printf("get login message: %v\n", err)
	}

	if login(info.UserID, info.Password) {
		fmt.Printf("login succeed！\n")
	} else {
		fmt.Printf("login failed！\n")
	}
}

func dealWithMessage(message commen.Message) {
	switch message.Type {
	case commen.LoginMessageType:
		dealWithLoginMessage(message.Data)
	case commen.ResponseMessageType:
		fmt.Println(commen.ResponseMessageType)
	default:
		fmt.Printf("other type")
	}
}

func dialogue(conn net.Conn) {
	defer conn.Close()

	// 循环的读取客户端的信息
	for {
		message, err := readDate(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client closed!\n")
				return
			}
			fmt.Printf("get login message error: %v", err)
		}
		dealWithMessage(message)
	}
}

func main() {
	fmt.Printf("服务端启动成功\n")

	listenr, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listenr.Close()
	if err != nil {
		fmt.Printf("some error when run server, error: %v", err)
	}

	for {
		fmt.Printf("等待客户端的连接......\n")

		conn, err := listenr.Accept()
		if err != nil {
			fmt.Printf("some error when accept server, error: %v", err)
		}

		// 一旦链接成功，在启动一个协程和客户端保持通讯
		go dialogue(conn)
	}
}
