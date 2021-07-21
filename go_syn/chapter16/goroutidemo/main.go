package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数 每隔1秒 输出hello world
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test()

	for i := 1; i < 10; i++ {
		fmt.Println("main hello go" + strconv.Itoa(i))
		//time.Sleep(time.Second)
	}

}
