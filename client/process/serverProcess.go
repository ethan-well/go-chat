package process

import (
	"errors"
	"fmt"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
)

func dealLoginResponse(responseMsg commen.ResponseMessage) (err error) {
	switch responseMsg.Code {
	case 200:
		err = nil
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

func dealRegisterResponse(responseMsg commen.ResponseMessage) (err error) {
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

// 处理服务端的返回
func Response(conn net.Conn, c chan bool, resErr chan error) (err error) {
	defer conn.Close()
	var responseMsg commen.ResponseMessage
	dispatcher := utils.Dispatcher{Conn: conn}

	responseMsg, err = dispatcher.ReadDate()
	if err != nil {
		fmt.Printf("some error, %v!\n", err)
		resErr <- err
		return
	}

	// 根据服务端返回的消息类型，进行相应的处理
	switch responseMsg.Type {
	case commen.LoginResponseMessageType:
		err = dealLoginResponse(responseMsg)
	case commen.RegisterResponseMessageType:
		err = dealRegisterResponse(responseMsg)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	resErr <- err
	c <- true

	return
}
