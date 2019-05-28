package main

import "fmt"

func main() {
	var key int
	var loop = true
	var userID int
	var password string

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

			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("输入用户密码:")
			fmt.Scanf("%s\n", &password)

			err := login(userID, password)
			if err != nil {
				fmt.Printf("登陆失败")
			} else {
				fmt.Printf("登陆成功")
			}
			loop = false
		case 2:
			fmt.Println("注册聊天室")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false // 等价 os.Exit(0)
		default:
			fmt.Printf("输入错误，请储物1、2、3\n")
		}
	}
}
