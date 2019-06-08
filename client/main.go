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
		fmt.Println("----------------欢迎使用多人聊天系统--------------")
		fmt.Println("\t\t请选择操作类型，选择 1、2、3")
		fmt.Println("\t\t\t 1、登陆")
		fmt.Println("\t\t\t 2、注册")
		fmt.Println("\t\t\t 3、退出")

		// 获取用户输入
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")

			fmt.Println("请输入用户名:")
			fmt.Scanf("%s\n", &userName)
			fmt.Println("输入用户密码:")
			fmt.Scanf("%s\n", &password)

			// err := login(userName, password)
			up := process.UserProcess{}
			err := up.Login(userName, password)

			if err != nil {
				fmt.Printf("Login failed: %v\n", err)
			} else {
				fmt.Printf("Login succeed!\n")
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
			fmt.Println("退出聊天室...")
			loop = false // 等价 os.Exit(0)
		default:
			fmt.Printf("输入错误，请储物1、2、3\n")
		}
	}
}
