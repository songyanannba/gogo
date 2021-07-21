package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int //指向队列队首
	tail    int //指向队尾
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

//判断环形队列为空
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func (this *CircleQueue) ShowQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempHand := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d] = %d \t", tempHand, this.array[tempHand])
		tempHand = (tempHand + 1) % this.maxSize
	}
	fmt.Println()
}

func main() {
	//先创建一个队列
	que := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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
			err := que.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("ok")
			}
		case "get":
			fmt.Println("get")
			val, err := que.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(val)
			}
		case "show":
			que.ShowQueue()

		case "exit":
			os.Exit(0)
		}
	}
}
