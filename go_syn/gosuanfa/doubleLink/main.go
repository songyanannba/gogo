package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nikename string
	next     *HeroNode
	pre      *HeroNode
}

//双向
func insert(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func del(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp = temp.next.pre
		}
	} else {
		fmt.Println("没有要删除掉的")
	}

}

func insert2(head *HeroNode, newHeroNode *HeroNode) {
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
		fmt.Println("不存在")
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

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
	}
}

func ListHeroNode2(head *HeroNode) {
	temp := head

	//是不是空
	if temp.next == nil {
		fmt.Println("空 list")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//遍历
	for {
		fmt.Printf("show list[%d ,%s ,%s]==>\n", temp.no, temp.name, temp.nikename)

		temp = temp.pre
		if temp.pre == nil { //表示找到最后
			break
		}
	}
}

func main() {

	head := &HeroNode{}

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

	hero5 := &HeroNode{
		no:       5,
		name:     "要改",
		nikename: "sdf",
	}

	insert2(head, hero1)
	insert2(head, hero3)
	insert2(head, hero5)
	insert2(head, hero2)

	ListHeroNode(head)
	fmt.Println("逆序")
	ListHeroNode2(head)

}
