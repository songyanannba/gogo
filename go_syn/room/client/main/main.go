package main

import (
	"fmt"
	"room/client/process"
)

//定义2个变量 1个表示用户ID 一个表示用户的密码
var userId int
var userPwd string
var userName string

func main() {

	//接收用户选项
	var key int

	var loop = true

	for loop {
		fmt.Println("\t\t---------欢迎登陆多人聊天系统---------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 推出系统")
		fmt.Println("\t\t\t 请选择(1--3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")

			//说明用户要登陆
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			//完成登陆
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字")
			fmt.Scanf("%s\n", &userName)

			//完成注册
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("推出系统")

		default:
			fmt.Println("输入有误 重新输入")
		}
	}

}
