package main

import "fmt"

type body struct {
	no   int
	next *body //指向下一个小孩的指针 默认值nil
}

//num 小孩的个数
func AddBody(num int) *body {
	first := &body{}

	curBoy := &body{}

	if num < 1 {
		fmt.Println("num no true")
		return first
	}

	//循环构建的连标
	for i := 1; i <= num; i++ {
		//因为第一个小孩比较特殊
		body := &body{
			no: i,
		}
		if i == 1 {
			first = body
			curBoy = body
			curBoy.next = first
		} else {
			curBoy.next = body
			curBoy = body
			curBoy.next = first
		}
	}
	return first
}

//
func list(first *body) {

	if first.next == nil {
		fmt.Println("没有人 连标为空")
		return
	}
	curboy := first
	for {
		fmt.Printf("编号 %d \n", curboy.no)
		if curboy.next == first {
			fmt.Println("遍历完")
			break
		}
		curboy = curboy.next
	}
}

func PlayGame(first *body, startNo int, countNum int) {
	if first.next == nil {
		fmt.Println("空 body")
		return
	}

	tail := first

	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}

	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("编号为%d 出圈\n", first.no)
		first = first.next
		tail.next = first

		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈编号为%d\n", first.no)

}

func main() {

	first := AddBody(5)
	list(first)
	PlayGame(first, 2, 3)

}
