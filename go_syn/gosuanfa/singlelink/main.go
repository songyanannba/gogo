package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nikename string
	next     *HeroNode
}

func del(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil { //表示找到最后
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("删除的id 不存在")
	}
	return
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //表示找到最后
			break
		} else if temp.next.no > newHeroNode.no {

			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("存在no")
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next
	}
	//将 newHeroNode加入到temp
	temp.next = newHeroNode
}

func ListHeroNode(head *HeroNode) {
	temp := head
	//是不是空
	if temp.next == nil {
		fmt.Println("空 list")
		return
	}
	//遍历
	for {
		fmt.Printf("show list[%d ,%s ,%s]==>\n", temp.next.no, temp.next.name, temp.next.nikename)

		temp = temp.next
		if temp.next == nil { //表示找到最后
			break
		}
		//time.Sleep(time.Second)
	}
}

func main() {

	//1 创建头节点
	head := &HeroNode{}

	//创建新的HerdNode节点
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nikename: "及时雨",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "武松",
		nikename: "大户",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "要改",
		nikename: "dd",
	}

	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	//InsertHeroNode2(head, hero5)

	ListHeroNode(head)

	del(head, 3)

	fmt.Println("+++")

	ListHeroNode(head)
}
