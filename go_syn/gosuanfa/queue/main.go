package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组模拟对列
	front   int    //表示指向队首
	rear    int    //表示指向队尾
}

//添加队列
func (this *Queue) AddQueue1(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = val

	return
}

//显示队列 找到队首 然后遍历到队尾
func (this *Queue) showQueue() {
	fmt.Println("队列当前情况是...")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\t", i, this.array[i])
	}
}

func (this *Queue) GetQueue() (val int, err error) {
	//判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

func main() {
	//先创建一个队列
	que := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	fmt.Printf("%v\n", que)

	var key string
	var val int
	for {
		fmt.Println("1 输入add 表示添加数据到队列")
		fmt.Println("1 输入get 表示获取数据")
		fmt.Println("1 输入show 表示展示数据")
		fmt.Println("1 输入exit 表示展示数据")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数据")
			fmt.Scanln(&val)
			err := que.AddQueue1(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("ok")
			}
		case "get":
			fmt.Println("get")
			val, err := que.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(val)
			}
		case "show":
			que.showQueue()

		case "exit":
			os.Exit(0)
		}
	}
}
