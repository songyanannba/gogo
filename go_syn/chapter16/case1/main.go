package main

import (
	"fmt"
	"sync"
)

//1编写一个函数 计算阶乘 并放入map
//2协程多个 统计结果 放入map
//3map 全局

var (
	myMap = make(map[int]int, 10)
	//声明全局变量 loke
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {

	for i := 1; i <= 12; i++ {
		go test(i)
	}

	//time.Sleep(time.Second * 2)
	//输出结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]%d\n", i, v)
	}
	lock.Unlock()

}
