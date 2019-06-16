package model

import (
	"errors"
)

//根据业务逻辑需要，自定义一些错误
var (
	ERROR_USER_DOES_NOT_EXIST = errors.New("User does not exist!")
	ERROR_USER_PWD            = errors.New("Password is invalid!")

	// status code for register
	ERROR_USER_ALREADY_EXISTS     = errors.New("Username already exists!")
	ERROR_PASSWORD_DOES_NOT_MATCH = errors.New("Password does not match!")
)
