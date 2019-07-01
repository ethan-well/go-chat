package process

import (
	"encoding/json"
	"fmt"
	common "go-chat/common/message"
	"go-chat/server/model"
	"go-chat/server/utils"
)

type GroupMessageProcess struct{}

// send messsage to all user in the target group
func (gmp GroupMessageProcess) sendToGroupUsers(message string) (err error) {
	// var info common.UserSendGroupMessage
	// err = json.Unmarshal([]byte(message), &info)
	var userSendGroupMessage common.UserSendGroupMessage
	err = json.Unmarshal([]byte(message), &userSendGroupMessage)
	if err != nil {
		fmt.Printf("some error when  json Unmarshal: %v\n", err)
	}

	// group message sender
	sourceUserName := userSendGroupMessage.UserName

	var toClientMessage common.ResponseMessage
	toClientMessage.Type = common.SendGroupMessageToClientType
	toClientMessage.Data = message

	data, err := json.Marshal(toClientMessage)
	if err != nil {
		fmt.Printf("json.Marshal(toClientMessage) error\n")
	}

	for id, connInfo := range model.ClientConnsMap {
		// do not send message to then sender
		if sourceUserName == connInfo.UserName {
			continue
		}

		fmt.Printf("client id: %v \n", id)

		dispatcher := utils.Dispatcher{Conn: connInfo.Conn}

		err = dispatcher.WriteData(data)
		if err != nil {
			fmt.Printf("conn err: %v\n", err)
		} else {
			fmt.Println("send succeed!")
		}
	}

	return
}
