package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name string
	Age  int
}

func chan0() {
	catChan := make(chan interface{}, 3)
	catChan <- 10
	catChan <- "tom jack"
	cat := Cat{"小猫", 3}
	catChan <- cat
	<-catChan
	<-catChan
	ncat := <-catChan
	fmt.Printf("ttt = %T  vvv == %v\n", ncat, ncat)
	a := ncat.(Cat)
	fmt.Printf("cat.name == %v \n", a.Name)
}

func chan1() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	//close(intChan)
	<-intChan
	<-intChan
	n := <-intChan
	fmt.Printf("ookkk %v", n)
}

func chan2() {
	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	close(intChan2)
	for v := range intChan2 {
		fmt.Printf("intChan2 == %v\n", v)
	}

}

func wri(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		time.Sleep(time.Second)
	}
	close(intChan)
}

func red(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("reader 数据==%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func chan3() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go wri(intChan)
	go red(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func putNum(intChan chan int) {
	for i := 0; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Printf("you一个 primeNum 协程取不到数据 退出\n")
	exitChan <- true
}

func chan4() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//开启协程 添加整数
	go putNum(intChan)

	//开启4个协程
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数有 %d \n ", res)
	}

	fmt.Printf("out")

}

func main() {

	//chan0()
	//chan1()
	//chan2()
	//chan3()
	chan4()
}
