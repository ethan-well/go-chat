package model

import (
	"errors"
)

//根据业务逻辑需要，自定义一些错误.
var (
	ERROR_USER_NOT_EXISTS = errors.New("user is not exists..")
	ERROR_USER_EXISTS     = errors.New("user has exists...")
	ERROR_USER_PWD        = errors.New("password is not valied")
)
