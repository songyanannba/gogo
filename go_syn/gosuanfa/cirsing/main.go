package main

import (
	"fmt"
)

type node struct {
	no   int
	name string
	next *node
}

func insert(head *node, newNode *node) {
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head
}

func list(head *node) {
	fmt.Println("打印链表")

	temp := head
	if temp.next == nil {
		fmt.Println("空")
		return
	}
	for {
		fmt.Printf("[no= %d name = %s ==>]\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func del(head *node, id int) *node {
	temp := head
	helper := head

	if temp.next == nil {
		fmt.Println("这是一个空的连标...")
		return head
	}

	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
	}

	//将heper
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true

	for {
		//fmt.Println("dfgsjkdh")
		if temp.next == head { //已经比较到最后一个
			break
		}
		if temp.no == id {
			if temp == head {
				//head = head.next
			}
			//删除
			helper.next = temp.next
			fmt.Printf("中间id == %d \n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}

	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("最后id == %d \n", id)
		} else {
			fmt.Printf("没有ID %d \n", id)
		}

	}
	return head
}

func main() {

	head := &node{}
	newNode1 := &node{
		no:   1,
		name: "张三",
	}

	newNode2 := &node{
		no:   2,
		name: "李四",
	}

	newNode3 := &node{
		no:   3,
		name: "呵呵",
	}

	newNode4 := &node{
		no:   4,
		name: "大的",
	}

	insert(head, newNode1)
	insert(head, newNode2)
	insert(head, newNode3)
	insert(head, newNode4)
	list(head)

	fmt.Println()

	del(head, 1)

	fmt.Println()
	fmt.Println()
	list(head)
}
