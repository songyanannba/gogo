package main

import (
	"flag"
	"fmt"
)

func main() {

	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认是hosthost")
	flag.IntVar(&port, "port", 3306, "默认为3306")

	flag.Parse()

	fmt.Printf("user=%v , pwd=%v host=%v port=%v", user, pwd, host, port)

}
