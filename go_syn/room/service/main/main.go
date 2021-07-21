package main

import (
	"fmt"
	"net"
	"room/service/model"
	"time"
)

func process(conn net.Conn) {
	//调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.pro2()
	if err != nil {
		fmt.Printf("processor.pro2 err %v \n", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//初始化链接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	//提示信息
	fmt.Println("服务器监听8889....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Printf(" net.Listen err =%v \n", err)
		return
	}

	for {
		fmt.Println("等待客户来链接...")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Printf("listen.Accept err %v \n", err)
		}
		//链接成功 启动一个协程 和客户端保持通信
		go process(conn)
	}

}
