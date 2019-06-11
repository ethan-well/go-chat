package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-chat/client/model"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
)

func dealLoginResponse(responseMsg commen.ResponseMessage) (err error) {
	switch responseMsg.Code {
	case 200:
		// 解析当前用户信息
		var userInfo commen.UserInfo
		err = json.Unmarshal([]byte(responseMsg.Data), &userInfo)
		if err != nil {
			return
		}

		// 初始化 CurrentUser
		user := model.User{}
		err = user.InitCurrentUser(userInfo.ID, userInfo.UserName)
		fmt.Printf("current user, id: %d, name: %v\n", model.CurrentUser.UserID, model.CurrentUser.UserName)
		if err != nil {
			return
		}
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
	if responseMsg.Code != 200 {
		err = errors.New("Server Error!")
		return
	}

	var userList []commen.UserInfo
	err = json.Unmarshal([]byte(responseMsg.Data), &userList)
	if err != nil {
		return
	}

	fmt.Println("On line user list")
	fmt.Printf("\t\tID\t\tname\n")
	for _, info := range userList {
		fmt.Printf("\t\t%v\t\t%v\n", info.ID, info.UserName)
	}
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
			if err != nil {
				fmt.Printf("some error when get online user info: %v\n", err)
			}
		default:
			fmt.Println("un")
		}
	}
}
