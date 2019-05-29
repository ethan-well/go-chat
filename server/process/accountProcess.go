package process

import (
	"encoding/json"
	commen "go-chat/commen/message"
	"go-chat/server/model"
)

func login(userName, passWord string) (user model.User, err error) {
	// 判断用户名和密码
	user, err = model.CurrentUserDao.Login(userName, passWord)
	return
}

func userLogin(message string) (code int, err error) {
	var info commen.LoginMessage
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		code = commen.ServerError
	}

	_, err = login(info.UserName, info.Password)
	switch err {
	case nil:
		code = commen.LoginSucceed
	case model.ERROR_USER_NOT_EXISTS:
		code = 404
	case model.ERROR_USER_PWD:
		code = 403
	default:
		code = 500
	}
	return
}
