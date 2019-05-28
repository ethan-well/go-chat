package process

import (
	"encoding/json"
	commen "go-chat/commen/message"
)

func login(userID int, passWord string) bool {
	// 判断用户名和密码
	return userID == 100 && passWord == "123"
}

func userLogin(message string) (code int, err error) {
	var info commen.LoginMessage
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		code = commen.ServerError
	}

	if login(info.UserID, info.Password) {
		code = commen.LoginSucceed
	} else {
		code = commen.LoginError
	}
	return
}
