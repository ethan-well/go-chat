package process

import (
	"encoding/json"
	"fmt"
	commen "go-chat/commen/message"
	"go-chat/server/model"
	"go-chat/server/utils"
	"net"
)

type OnlineInfoProcess struct {
	Conn net.Conn
}

type UserInfo = commen.UserInfo

func (this OnlineInfoProcess) showAllOnlineUserList() (err error) {
	var onlineUserList []UserInfo
	var code int
	for _, connInfo := range model.ClientConnsMap {
		user, err := model.CurrentUserDao.GetUsrByUserName(connInfo.UserName)
		if err != nil {
			continue
		}
		userInfo := UserInfo{user.ID, user.Name}
		onlineUserList = append(onlineUserList, userInfo)
	}

	data, err := json.Marshal(onlineUserList)

	if err != nil {
		code = commen.ServerError
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
	responseMessage := commen.ResponseMessage{
		Code:  code,
		Type:  commen.ShowAllOnlineUsersType,
		Data:  data,
		Error: errMsg,
	}

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Printf("some error when generate response message, error: %v", err)
	}

	dispatcher := utils.Dispatcher{Conn: conn}

	err = dispatcher.WirteData(responseData)
	if err != nil {
		fmt.Printf("some error: %v", err)
	}
	return
}
