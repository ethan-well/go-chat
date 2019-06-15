package main

import (
	"fmt"
	"go-chat/client/logger"
	"go-chat/client/process"
)

func main() {
	var (
		key              int
		loop             = true
		userName         string
		password         string
		password_confirm string
	)

	for loop {
		logger.Info("\n----------------Welcome to the chat room--------------\n")
		logger.Info("\t\tSelect the options：\n")
		logger.Info("\t\t\t 1、Sign in\n")
		logger.Info("\t\t\t 2、Sign up\n")
		logger.Info("\t\t\t 3、Exit the system\n")

		// get user input
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			logger.Info("sign In Please\r\n")
			logger.Notice("Username:\n")
			fmt.Scanf("%s\n", &userName)
			logger.Notice("Password:\n")
			fmt.Scanf("%s\n", &password)

			// err := login(userName, password)
			up := process.UserProcess{}
			err := up.Login(userName, password)

			if err != nil {
				logger.Error("Login failed: %v\r\n", err)
			} else {
				logger.Success("Login succeed!\r\n")
			}
		case 2:
			logger.Info("Create account\n")
			logger.Notice("user name：\n")
			fmt.Scanf("%s\n", &userName)
			logger.Notice("password：\n")
			fmt.Scanf("%s\n", &password)
			logger.Notice("password confirm：\n")
			fmt.Scanf("%s\n", &password_confirm)

			up := process.UserProcess{}
			err := up.Register(userName, password, password_confirm)
			if err != nil {
				logger.Error("Creae account failed: %v\n", err)
			}
		case 3:
			logger.Warn("Exit...\n")
			loop = false // this is equal to 'os.Exit(0)'
		default:
			logger.Error("Select is invalid!\n")
		}
	}
}
