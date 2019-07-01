package process

import (
	"encoding/json"
	"fmt"
	common "go-chat/common/message"
	"go-chat/server/model"
	"go-chat/server/utils"
	"net"
)

type OnlineInfoProcess struct {
	Conn net.Conn
}

type UserInfo = common.UserInfo

func (this OnlineInfoProcess) showAllOnlineUserList() (err error) {
	var onlineUserList []UserInfo
	var code int
	for _, connInfo := range model.ClientConnsMap {
		user, err := model.CurrentUserDao.GetUserByUserName(connInfo.UserName)
		if err != nil {
			continue
		}
		userInfo := UserInfo{ID: user.ID, UserName: user.Name}
		onlineUserList = append(onlineUserList, userInfo)
	}

	data, err := json.Marshal(onlineUserList)

	if err != nil {
		code = common.ServerError
	} else {
		code = 200
	}

	err = responseClient(this.Conn, code, string(data), fmt.Sprintf("%v", err))
	if err != nil {
		fmt.Printf("point to point communicate, response client error: %v", err)
	}
	return
}

func responseClient(conn net.Conn, code int, data string, errMsg string) (err error) {
	responseMessage := common.ResponseMessage{
		Code:  code,
		Type:  common.ShowAllOnlineUsersType,
		Data:  data,
		Error: errMsg,
	}

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Printf("some error when generate response message, error: %v", err)
	}

	dispatcher := utils.Dispatcher{Conn: conn}

	err = dispatcher.WriteData(responseData)
	if err != nil {
		fmt.Printf("some error: %v", err)
	}
	return
}
