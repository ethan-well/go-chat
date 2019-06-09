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
	if err != nil {
		return
	}

	// get response from server
	// get_response
	return
}
