package model

import (
	"errors"
)

//根据业务逻辑需要，自定义一些错误
var (
	ERROR_USER_NOT_EXISTS = errors.New("user is not exists..")
	ERROR_USER_PWD        = errors.New("password is not valied")

	// status code for register
	ERROR_USER_EXISTED       = errors.New("user name is already existed!")
	ERROR_PASSWORD_NOT_MATCH = errors.New("passworld not match!")
)
