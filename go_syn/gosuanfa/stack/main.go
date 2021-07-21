package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //栈最大数量
	Top    int //栈顶
	arr    [5]int
}

func (this *Stack) Push(val int) error {

	if this.MaxTop-1 == this.Top {
		fmt.Println("栈满...")
		return errors.New("stack full")
	}

	this.Top++
	this.arr[this.Top] = val
	return nil
}

func (this *Stack) List() {

	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	for i := 0; i <= this.Top; i++ {
		fmt.Printf("val %v\n", this.arr[i])
	}
}

func (this *Stack) Pop() (val int, err error) {
	//判断是否为空
	if this.Top == -1 {
		return 0, errors.New("stack empty")
	}

	val = this.arr[this.Top]
	this.Top--
	return val, nil

}

func main() {

	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Printf("main stack %v= \n", stack)
	stack.List()

	val, _ := stack.Pop()
	fmt.Printf("pop val %v\n", val)

	stack.List()

	val, _ = stack.Pop()
	fmt.Printf("pop val %v\n", val)

	stack.List()
	val, _ = stack.Pop()
	fmt.Printf("pop val %v\n", val)

	stack.List()
}
