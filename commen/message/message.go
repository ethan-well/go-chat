package commen

const (
	LoginMessageType    = "LoginMessage"
	ResponseMessageType = "ResponseMessage"
	RegisterMessageType = "RegisterMessage"

	ServerError = 500

	// status code for login
	LoginError   = 403
	NotExit      = 404
	LoginSucceed = 200

	// status code for register
	HaveExisted       = 403
	RegisterSucceed   = 200
	PassworldNotMatch = 402
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserName string
	Password string
}

type ResponseMessage struct {
	Code  int    // 404 用户没找到， 403 账号或者密码错误, 200 登陆成功, 500 服务端错误
	Error string // 错误消息
}

type RegisterMessage struct {
	UserName        string
	Password        string
	PasswordConfirm string
}
