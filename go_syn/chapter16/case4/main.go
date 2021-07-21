package main

import (
	"fmt"
)

func sel() {
	intchan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intchan <- i
	}

	stringchan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	//遍历如果没有关闭的管道 会报错 deadlock
	//在实际开发中 可能不好确定什么时候关闭管道
	//可以用select
	for {
		select {
		case v := <-intchan:
			fmt.Printf("intchan 数据 %d \n", v)
			//time.Sleep(time.Second)
		case s := <-stringchan:
			fmt.Printf("stringchan 数据 %v \n", s)
			//time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到数据 不玩了 \n")
			return
		}
	}

}

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("h w %d\n", i)
		//time.Sleep(time.Second)
	}
}

func f2() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("f2 err %v\n", err)
		}
	}()

	var myMap map[int]string
	//myMap = make(map[int]string, 10)
	myMap[0] = "sdjfgh"
}

func main() {
	//sel()

	go f1()
	go f2()
	for i := 0; i < 10; i++ {
		fmt.Printf("ok %d\n", i)
	}
}
