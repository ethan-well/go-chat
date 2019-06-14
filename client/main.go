package main

import (
	"fmt"
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
		fmt.Println("----------------Welcome to the chat room--------------")
		fmt.Println("\t\tSelect the options：")
		fmt.Println("\t\t\t 1、Sign in")
		fmt.Println("\t\t\t 2、Sign up")
		fmt.Println("\t\t\t 3、Exit the system")

		// 获取用户输入
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Printf("sign In\r\n")
			fmt.Println("Username:")
			fmt.Scanf("%s\n", &userName)
			fmt.Println("Password:")
			fmt.Scanf("%s\n", &password)

			// err := login(userName, password)
			up := process.UserProcess{}
			err := up.Login(userName, password)

			if err != nil {
				fmt.Printf("Login failed: %v\r\n", err)
			} else {
				fmt.Printf("Login succeed!\r\n")
			}
		case 2:
			fmt.Println("Create account")
			fmt.Println("user name：")
			fmt.Scanf("%s\n", &userName)
			fmt.Println("password：")
			fmt.Scanf("%s\n", &password)
			fmt.Println("password confirm：")
			fmt.Scanf("%s\n", &password_confirm)

			up := process.UserProcess{}
			err := up.Register(userName, password, password_confirm)
			if err != nil {
				fmt.Printf("Creae account failed")
			}
			loop = false
		case 3:
			fmt.Println("Exit...")
			loop = false // 等价 os.Exit(0)
		default:
			fmt.Printf("Select is invalid!\n")
		}
	}
}
