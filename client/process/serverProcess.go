package process

import (
	"encoding/json"
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

func dealGroupMessage(responseMsg commen.ResponseMessage) (err error) {
	var groupMessage commen.SendGroupMessageToClient
	err = json.Unmarshal([]byte(responseMsg.Data), &groupMessage)
	if err != nil {
		return
	}
	fmt.Printf("%v send message: %v\n", groupMessage.UserID, groupMessage.Content)
	return
}

func showAllOnlineUsersList(responseMsg commen.ResponseMessage) (err error) {
	fmt.Println("deal with show all online users")
	return
}

// 处理服务端的返回
func Response(conn net.Conn) (err error) {
	var responseMsg commen.ResponseMessage
	dispatcher := utils.Dispatcher{Conn: conn}

	for {
		responseMsg, err = dispatcher.ReadDate()
		if err != nil {
			fmt.Printf("some error, %v!\n", err)
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
		case commen.SendGroupMessageToClientType:
			err = dealGroupMessage(responseMsg)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case commen.ShowAllOnlineUsersType:
			err = showAllOnlineUsersList(responseMsg)
		default:
			fmt.Println("un")
		}
	}
}
