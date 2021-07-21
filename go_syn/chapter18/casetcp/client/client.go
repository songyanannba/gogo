package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "10.177.112.53:8888")

	if err != nil {
		fmt.Printf("链接失败 err %v\n", err)
	}

	fmt.Printf("conn 成功...%v\n", conn)

	//功能1 客户端输入数据
	reader := bufio.NewReader(os.Stdin) //os.stdin 标准输入 终端

	for {
		//从终端读取数据
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("终端读取数据错误 ReadString", err)
		}

		str = strings.Trim(str, "\r\n")
		if str == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		//数据 发送给服务器
		_, err = conn.Write([]byte(str + "\n"))
		if err != nil {
			fmt.Println("conn.Write err", err)
		}

	}

}
