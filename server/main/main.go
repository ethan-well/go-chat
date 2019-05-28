package main

import (
	"encoding/json"
	"fmt"
	commen "go-chat/commen/message"
	"go-chat/server/utils"
	"io"
	"net"
)

func login(userID int, passWord string) bool {
	// 判断用户名和密码
	return userID == 100 && passWord == "123"
}

func dealWithLoginMessage(message string) (code int, err error) {
	var info commen.LoginMessage
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		code = commen.ServerError
	}

	if login(info.UserID, info.Password) {
		code = commen.LoginSucceed
	} else {
		code = commen.LoginError
	}
	return
}

func responseClient(conn net.Conn, code int, err error) {
	var responseMessage commen.ResponseMessage
	responseMessage.Code = code
	if err != nil {
		responseMessage.Error = fmt.Sprintf("login error: %v", err)
	}

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Printf("some error when generate response message, error: %v", err)
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.WirteData(responseData)
}

func dialogue(conn net.Conn) {
	defer conn.Close()

	// 循环的读取客户端的信息
	dispatcher := utils.Dispatcher{Conn: conn}
	for {
		message, err := dispatcher.ReadData()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client closed!\n")
				return
			}
			fmt.Printf("get login message error: %v", err)
		}
		// code, err := dealWithMessage(message)
		code, err := process(message)
		// 返回状态码给客户端

		responseClient(conn, code, err)
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
		fmt.Printf("等待客户端的连接...\n")

		conn, err := listenr.Accept()
		if err != nil {
			fmt.Printf("some error when accept server, error: %v", err)
		}

		// 一旦链接成功，在启动一个协程和客户端保持通讯
		go dialogue(conn)
	}
}
