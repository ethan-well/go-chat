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

	fmt.Println("read data from client!")
	// 读取小心长度信息
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

	err = json.Unmarshal(buf[:dataLen], &msg)
	if err != nil {
		fmt.Printf("json.Unmarshl error: %v", err)
	}

	fmt.Printf("login message: %v", msg)
	return
}

func dialogue(conn net.Conn) {
	defer conn.Close()

	// 循环的读取客户端的信息
	for {
		_, err := readDate(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client closed!\n")
				return
			}
			fmt.Printf("get login message error: %v", err)
		}
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
