package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//循环接受 客户端发送数据
	defer conn.Close()
	for {
		//创建切片
		buf := make([]byte, 1024)
		//
		//fmt.Printf("等待客户的输入 %s \n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			//fmt.Printf("服务器 conn.Read() err = %v \n", err)
			fmt.Println("客户端已经推出")
			return
		}
		//显示客户发送的内容
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("开始监听服务器...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败...")
		return
	}

	defer listen.Close()
	fmt.Printf("lesten %v \n", listen)

	//循环等待 客户端 链接
	for {
		fmt.Printf("等待客户端的链接...")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Printf("Accept() err = %v \n", err)
		} else {
			fmt.Printf("accept() suc com = %v 客户端ip = %v \n", conn, conn.RemoteAddr().String())
		}
		//准备协程 为客户端提供服务
		fmt.Printf("准备协程 为客户端提供服务...\n")
		go process(conn)
	}

}
