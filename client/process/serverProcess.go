package process

import (
	"errors"
	"fmt"
	"go-chat/client/utils"
	commen "go-chat/commen/message"
	"net"
)

func Response(conn net.Conn) (err error) {
	defer conn.Close()
	var responseMsg commen.ResponseMessage
	dispatcher := utils.Dispatcher{Conn: conn}

	responseMsg, err = dispatcher.ReadDate()
	if err != nil {
		fmt.Printf("some error, %v!\n", err)
		return
	}

	switch responseMsg.Code {
	case 200:
		fmt.Printf("Loggin succeed!")
	case 500:
		err = errors.New("server error")
	case 404:
		err = errors.New("user not exist")
	case 403:
		err = errors.New("pasword not valide")
	default:
		err = errors.New("some error")
	}
	return
}
