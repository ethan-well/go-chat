package commen

const (
	LoginMessageType    = "LoginMessage"
	ResponseMessageType = "ResponseMessage"
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
	Code  int    // 404 用户没找到， 403 账号或者密码错误, 200 登陆成功
	Error string // 错误消息
}
