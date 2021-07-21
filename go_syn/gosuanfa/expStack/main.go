package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int //栈最大数量
	Top    int //栈顶
	arr    [20]int
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

func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

func (this *Stack) cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	}
	return res
}

func (this *Stack) Priority(oper int) int {
	op := 0
	if oper == 42 || oper == 47 {
		op = 1
	} else if oper == 43 || oper == 45 {
		op = 0
	}
	return op
}

func main() {

	//通过栈计算加减乘除
	numstack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	operstack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	expstr := "30+20*5-32"

	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	kepNum := ""

	for {
		ch := expstr[index : index+1]
		temp := int([]byte(ch)[0])
		if operstack.IsOper(temp) {
			if operstack.Top == -1 { //空
				operstack.Push(temp)
			} else {
				if operstack.Priority(operstack.arr[operstack.Top]) >= operstack.Priority(temp) {
					num1, _ = numstack.Pop()
					num2, _ = numstack.Pop()
					oper, _ = operstack.Pop()
					result = operstack.cal(num1, num2, oper)
					numstack.Push(result)
					operstack.Push(temp)
				} else {
					operstack.Push(temp)
				}
			}
		} else {
			kepNum += ch
			if index == len(expstr)-1 {
				val, _ := strconv.ParseInt(kepNum, 10, 64)
				numstack.Push(int(val))
			} else {
				if operstack.IsOper(int([]byte(expstr[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(kepNum, 10, 64)
					numstack.Push(int(val))
					kepNum = ""
				}
			}

		}
		if index+1 == len(expstr) {
			break
		}
		index++
	}

	for {
		if operstack.Top == -1 { //空
			break
		}
		num1, _ = numstack.Pop()
		num2, _ = numstack.Pop()
		oper, _ = operstack.Pop()
		result = operstack.cal(num1, num2, oper)
		numstack.Push(result)
	}

	//
	v, _ := numstack.Pop()
	fmt.Println(expstr+"=", v)

}
