package main

import (
	"fmt"
	"go-chat/server/model"
	"go-chat/server/process"
	"net"
	"time"
)

func init() {
	// 初始化 redis 连接池，全局唯一
	initRedisPool(16, 0, time.Second*300, "127.0.0.1:6379")

	// 创建 userDao 用于操作用户信息
	// 全局唯一 UserDao 实例：model.CurrentUserDao
	model.CurrentUserDao = model.InitUserDao(pool)
}

// 和客户端的通信交互
// conn 就是客户端和服务器之间建立的连接
// 每当有个用户登陆进来之后，就启动一个 go routine
// 这个 go routine 专门用来处理服务器和客户端的通信
func dialogue(conn net.Conn) {
	defer conn.Close()
	processor := process.Processor{Conn: conn}
	processor.MainProcess()
}

func main() {
	fmt.Printf("Server is already\n")

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listener.Close()
	if err != nil {
		fmt.Printf("some error when run server, error: %v", err)
	}

	for {
		fmt.Printf("Waiting for client...\n")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("some error when accept server, error: %v", err)
		}

		// 一旦链接成功，在启动一个协程和客户端保持通讯
		go dialogue(conn)
	}
}
