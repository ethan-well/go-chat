package process

import (
	"encoding/json"
	"go-chat/client/utils"
	common "go-chat/common/message"
	"go-chat/config"
	"net"
)

type MessageProcess struct{}

// user send message to server
func (msgProc MessageProcess) SendGroupMessageToServer(groupID int, userName string, content string) (err error) {
	// connect server
	serverInfo := config.Configuration.ServerInfo
	conn, err := net.Dial("tcp", serverInfo.Host)

	if err != nil {
		return
	}

	var message common.Message
	message.Type = common.UserSendGroupMessageType

	// group message
	userSendGroupMessage := common.UserSendGroupMessage{
		GroupID:  groupID,
		UserName: userName,
		Content:  content,
	}
	data, err := json.Marshal(userSendGroupMessage)
	if err != nil {
		return
	}

	message.Data = string(data)
	data, _ = json.Marshal(message)

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)

	return
}

// request all online user
func (msg MessageProcess) GetOnlineUerList() (err error) {
	serverInfo := config.Configuration.ServerInfo
	conn, err := net.Dial("tcp", serverInfo.Host)
	if err != nil {
		return
	}

	var message = common.Message{}
	message.Type = common.ShowAllOnlineUsersType

	requestBody, err := json.Marshal("")
	if err != nil {
		return
	}
	message.Data = string(requestBody)

	data, err := json.Marshal(message)
	if err != nil {
		return
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)
	if err != nil {
		return
	}

	errMsg := make(chan error)
	go Response(conn, errMsg)
	err = <-errMsg
	if err != nil {
		return
	}

	for {
		showAfterLoginMenu()
	}

	return
}

func (msgProc MessageProcess) PointToPointCommunication(targetUserName, sourceUserName, message string) (conn net.Conn, err error) {
	serverInfo := config.Configuration.ServerInfo
	conn, err = net.Dial("tcp", serverInfo.Host)
	if err != nil {
		return
	}
	// defer conn.Close()

	var pointToPointMessage common.Message

	pointToPointMessage.Type = common.PointToPointMessageType

	messageBody := common.PointToPointMessage{
		SourceUserName: sourceUserName,
		TargetUserName: targetUserName,
		Content:        message,
	}

	data, err := json.Marshal(messageBody)
	if err != nil {
		return
	}

	pointToPointMessage.Data = string(data)

	data, err = json.Marshal(pointToPointMessage)
	if err != nil {
		return
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)
	return
}
