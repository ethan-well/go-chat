package main

import (
	"fmt"
	"net"
)

func dialogue(conn net.Conn) {
	defer conn.Close()

	// 循环的读取客户端的信息
	for {
		buf := make([]byte, 10240)
		fmt.Println("read data from client!")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Printf("n: %v\n", n)
			fmt.Printf("conn read data error: %v\n", err)
			return
		}

		fmt.Printf("data from client: %v\n", buf[0:4])
		return
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
