package process

import (
	"encoding/json"
	"fmt"
	commen "go-chat/commen/message"
	"go-chat/server/model"
	"go-chat/server/utils"
)

type GroupMessageProcess struct{}

// 向组内人员发送消息
func (gmp GroupMessageProcess) sendToGroupUsers(message string) (err error) {
	// var info commen.UserSendGroupMessage
	// err = json.Unmarshal([]byte(message), &info)
	var toClientMessage commen.ResponseMessage
	toClientMessage.Type = commen.SendGroupMessageToClientType
	toClientMessage.Data = message

	data, err := json.Marshal(toClientMessage)
	if err != nil {
		fmt.Printf("json.Marshal(toClientMessage) error\n")
	}

	for id, connInfo := range model.ClientConnsMap {
		fmt.Printf("client id: %v \n", id)

		dispatcher := utils.Dispatcher{Conn: connInfo.Conn}

		err = dispatcher.WirteData(data)
		if err != nil {
			fmt.Printf("conn err: %v\n", err)
		} else {
			fmt.Println("send succeed!")
		}
	}

	return
}
