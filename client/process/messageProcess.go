package process

import (
	"encoding/json"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
)

type MessageProcess struct{}

// user send message to server
func (msgProc MessageProcess) SendGroupMessageToServer(groupID, userID int, content string) (err error) {
	// connect server
	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		return
	}

	var message commen.Message
	message.Type = commen.UserSendGroupMessageType

	// group message
	userSendGroupMessage := commen.UserSendGroupMessage{
		GroupID: groupID,
		UserID:  userID,
		Content: content,
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
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		return
	}

	var message = commen.Message{}
	message.Type = commen.ShowAllOnlineUsersType

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

	go Response(conn)

	for {
		showAfterLoginMenu()
	}

	return
}

func (msgProc MessageProcess) PointToPointCommunication(targetUserName, sourceUserName, message string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		return
	}
	// defer conn.Close()

	var pointToPointMessage commen.Message

	pointToPointMessage.Type = commen.PointToPointMessageType

	messageBody := commen.PointToPointMessage{
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
