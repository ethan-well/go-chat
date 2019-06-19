package process

import (
	"encoding/json"
	"errors"
	"go-chat/client/logger"
	"go-chat/client/model"
	"go-chat/client/utils"
	common "go-chat/common/message"
	"net"
)

func dealLoginResponse(responseMsg common.ResponseMessage) (err error) {
	switch responseMsg.Code {
	case 200:
		// 解析当前用户信息
		var userInfo common.UserInfo
		err = json.Unmarshal([]byte(responseMsg.Data), &userInfo)
		if err != nil {
			return
		}

		// 初始化 CurrentUser
		user := model.User{}
		err = user.InitCurrentUser(userInfo.ID, userInfo.UserName)
		logger.Success("Login succeed!\n")
		logger.Notice("Current user, id: %d, name: %v\n", model.CurrentUser.UserID, model.CurrentUser.UserName)
		if err != nil {
			return
		}
	case 500:
		err = errors.New("Server error!")
	case 404:
		err = errors.New("User does not exist!")
	case 403:
		err = errors.New("Password invalid!")
	default:
		err = errors.New("Some error!")
	}
	return
}

func dealRegisterResponse(responseMsg common.ResponseMessage) (err error) {
	switch responseMsg.Code {
	case 200:
		logger.Success("Register succeed!\n")
	case 500:
		err = errors.New("Server error!")
	case 403:
		err = errors.New("User already exists!")
	case 402:
		err = errors.New("Password invalid!")
	default:
		err = errors.New("Some error!")
	}
	return
}

func dealGroupMessage(responseMsg common.ResponseMessage) (err error) {
	var groupMessage common.SendGroupMessageToClient
	err = json.Unmarshal([]byte(responseMsg.Data), &groupMessage)
	if err != nil {
		return
	}
	logger.Info("%v send you:", groupMessage.UserName)
	logger.Notice("\t%v\n", groupMessage.Content)
	return
}

func showAllOnlineUsersList(responseMsg common.ResponseMessage) (err error) {
	if responseMsg.Code != 200 {
		err = errors.New("Server Error!")
		return
	}

	var userList []common.UserInfo
	err = json.Unmarshal([]byte(responseMsg.Data), &userList)
	if err != nil {
		return
	}

	logger.Success("Online user list(%v users)\n", len(userList))
	logger.Notice("\t\tID\t\tname\n")
	for _, info := range userList {
		logger.Success("\t\t%v\t\t%v\n", info.ID, info.UserName)
	}

	return
}

func showPointToPointMessage(responseMsg common.ResponseMessage) (err error) {
	if responseMsg.Code != 200 {
		err = errors.New(responseMsg.Error)
		return
	}

	var pointToPointMessage common.PointToPointMessage
	err = json.Unmarshal([]byte(responseMsg.Data), &pointToPointMessage)
	if err != nil {
		return
	}

	logger.Info("\r\n\r\n%v say: ", pointToPointMessage.SourceUserName)
	logger.Notice("\t%v\n", pointToPointMessage.Content)
	return
}

// 处理服务端的返回
func Response(conn net.Conn, errMsg chan error) (err error) {
	var responseMsg common.ResponseMessage
	dispatcher := utils.Dispatcher{Conn: conn}

	for {
		responseMsg, err = dispatcher.ReadData()
		if err != nil {
			logger.Error("Waiting response error: %v\n", err)
			return
		}

		// 根据服务端返回的消息类型，进行相应的处理
		switch responseMsg.Type {
		case common.LoginResponseMessageType:
			err = dealLoginResponse(responseMsg)
			errMsg <- err
		case common.RegisterResponseMessageType:
			err = dealRegisterResponse(responseMsg)
			errMsg <- err
		case common.SendGroupMessageToClientType:
			err = dealGroupMessage(responseMsg)
			if err != nil {
				logger.Error("%v\n", err)
			}
		case common.ShowAllOnlineUsersType:
			err = showAllOnlineUsersList(responseMsg)
			errMsg <- err
		case common.PointToPointMessageType:
			err = showPointToPointMessage(responseMsg)
			errMsg <- err
		default:
			logger.Error("Unknown message type!")
		}

		if err != nil {
			return
		}
	}
}
