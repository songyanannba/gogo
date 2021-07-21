package main

import (
	"fmt"
	"io"
	"net"
	"room/commom/message"
	process2 "room/service/process"
	"room/service/utils"
)

type Processor struct {
	Conn net.Conn
}

//编写一个函数 根据客户端的消息类型 决定调用不同的函数
func (this *Processor) servicerProcessMes(mes *message.Message) (err error) {

	//根据类型
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆的逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		up.ServiceProcesslogin(mes)
	case message.RegisterMseType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		up.ServiceProcessRegister(mes)
	case message.SmsMesType:
		//处理注册
		fmt.Println("群聊")
		smsUp := &process2.SmsProcess{}

		smsUp.SendGroupMes(mes)

	default:
		fmt.Println("消息类型不存在 无法处理...")
	}

	return
}

func (this *Processor) pro2() (err error) {
	//循环读取客户端 发送的信息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Printf("readPkg err %v \n", err)
			if err == io.EOF {
				fmt.Println("tf.ReadPkg", err)
				return err
			} else {
				fmt.Println("错误")
				return err
			}
		}
		fmt.Printf("readPkg mes %v \n", mes)

		err = this.servicerProcessMes(&mes)
		if err != nil {
			fmt.Printf(" ServicerProcessMes err =%v \n", err)
			return err
		}
	}
}
