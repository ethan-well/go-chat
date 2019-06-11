package commen

const (
	LoginMessageType             = "LoginMessage"
	RegisterMessageType          = "RegisterMessage"
	LoginResponseMessageType     = "LoginResponseMessageType"
	RegisterResponseMessageType  = "ResponseMessageType"
	UserSendGroupMessageType     = "UserSendGroupMessageType"
	SendGroupMessageToClientType = "SendGroupMessageToClientType"
	ShowAllOnlineUsersType       = "ShowAllOnlineUsersType"
	PointToPointMessageType      = "PointToPointMessageType"

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
	Type  string
	Code  int    // 404 用户没找到， 403 账号或者密码错误, 200 登陆成功, 500 服务端错误
	Error string // 错误消息
	Data  string
}

type RegisterMessage struct {
	UserName        string
	Password        string
	PasswordConfirm string
}

type UserSendGroupMessage struct {
	GroupID int    // 目标组 id
	UserID  int    // 当前用户 id
	Content string // 小心主体内容
}

type SendGroupMessageToClient struct {
	GroupID int
	UserID  int
	Content string
}

type PointToPointMessage struct {
	SourceUserID   int
	SourceUserName string
	TargetUserID   int
	TargetUserName string
	Content        string
}

// on line user info
type UserInfo struct {
	ID       int
	UserName string
}
