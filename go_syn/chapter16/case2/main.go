package main

import "fmt"

func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("dfgdfg  %v\n", intChan)

	//向管道写入值
	intChan <- 10
	intChan <- 121
	intChan <- 12

	//管道的长度 管道的容量
	fmt.Printf("len = %v\n", len(intChan))
	fmt.Printf("cap = %v \n", cap(intChan))

	var mm int
	mm = <-intChan

	fmt.Printf("mm %d\n", mm)

	fmt.Printf("len = %v\n", len(intChan))
	fmt.Printf("cap = %v \n", cap(intChan))

}
