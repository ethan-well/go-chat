package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"go-chat/client/logger"
	common "go-chat/common/message"
	"net"
)

type Dispatcher struct {
	Conn net.Conn
	Buf  [10240]byte
}

func (dispatcher Dispatcher) ReadData() (msg common.ResponseMessage, err error) {
	buf := make([]byte, 10240)

	// 读取消息长度信息
	n, err := dispatcher.Conn.Read(buf[:4])
	if err != nil {
		return
	}
	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(buf[0:4])

	// 读取消息本身
	n, err = dispatcher.Conn.Read(buf[:dataLen])
	if err != nil {
		logger.Error("server read data login data error: %v", err)
	}

	// 对比消息本身的长度和期望长度是否匹配
	if n != int(dataLen) {
		err = errors.New("login message length error")
		return
	}

	// 从 conn 中解析消息并存放到 msg 中，此处一定传递的是 msg 的地址
	err = json.Unmarshal(buf[:dataLen], &msg)
	if err != nil {
		logger.Error("json.Unmarshal error: %v", err)
	}
	return
}

func (dispatcher Dispatcher) SendData(data []byte) (err error) {
	// 首先发送数据 data 的长度到服务器端
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)

	// 客户端发送消息长度
	writeLen, err := dispatcher.Conn.Write(bytes[:])
	if writeLen != 4 || err != nil {
		logger.Error("send data length to server error: %v", err)
		return
	}

	//客户端发送消息本身
	writeLen, err = dispatcher.Conn.Write(data)
	if err != nil {
		logger.Error("send data to server error: %v", err)
		return
	}
	return
}
