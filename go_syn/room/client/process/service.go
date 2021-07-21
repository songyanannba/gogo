package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"room/client/utils"
	"room/commom/message"
)

//显示登陆成功后的界面..
func ShowMenu() {
	fmt.Println("\t\t---------恭喜xxx登陆成功---------")
	fmt.Println("\t\t\t 1 显示用户列表")
	fmt.Println("\t\t\t 2 发送消息")
	fmt.Println("\t\t\t 3 消息列表")
	fmt.Println("\t\t\t 4 退出系统")
	fmt.Println("\t\t\t 请选择(1--4):")

	var key int
	fmt.Scanf("%d\n", &key)

	var content string

	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		fmt.Println("显示用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息:请输入你对大家说的话")
		fmt.Scanf("%s\n", &content)
		smsProcess.SenGroupMes(content)

	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("推出系统")
		os.Exit(0)
	default:
		fmt.Println("输入选项有误 重新输入")
	}
}

//和服务器保持连接
func serviceProcessMes(conn net.Conn) {
	//创建一个transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		fmt.Printf("客户端正在等待读取服务器发送的消息")
		if err != nil {
			fmt.Printf("tf.ReadPkg() err %v \n", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//处理
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Date), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsMesType:
			//处理

			outputGroupMes(&mes)

		default:
			fmt.Println("服务器返回未知消息类型")
		}
	}

}
