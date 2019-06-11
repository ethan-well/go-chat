package process

import (
	"encoding/json"
	commen "go-chat/commen/message"
	"go-chat/server/model"
	"go-chat/server/utils"
)

type PointToPointMessageProcess struct{}

func (this PointToPointMessageProcess) sendMessageToTargetUser(message string) (err error) {
	var pointToPointMessage commen.PointToPointMessage
	err = json.Unmarshal([]byte(message), &pointToPointMessage)
	if err != nil {
		return
	}

	clientConn := model.ClientConn{}
	conn, err := clientConn.SearchByUserName(pointToPointMessage.TargetUserName)
	if err != nil {
		return
	}

	var responseMessage commen.ResponseMessage
	responseMessage.Type = commen.PointToPointMessageType

	var responseMessageData = commen.PointToPointMessage{
		SourceUserName: pointToPointMessage.SourceUserName,
		TargetUserName: pointToPointMessage.TargetUserName,
		Content:        pointToPointMessage.Content,
	}

	data, err := json.Marshal(responseMessageData)
	if err != nil {
		return
	}
	responseMessage.Data = string(data)

	responseMessage.Code = 200

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		return
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.WirteData(responseData)

	return
}
