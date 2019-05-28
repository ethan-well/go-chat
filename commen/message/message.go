package commen

const (
	LoginMessageType    = "LoginMessage"
	ResponseMessageType = "ResponseMessage"

	ServerError  = 500
	LoginError   = 403
	NotExit      = 404
	LoginSucceed = 200
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserID   int
	Password string
}

type ResponseMessage struct {
	Code  int    // 404 用户没找到， 403 账号或者密码错误, 200 登陆成功, 500 服务端错误
	Error string // 错误消息
}
