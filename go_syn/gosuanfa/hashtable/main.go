package main

import (
	"fmt"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) InsertE(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil
	if cur == nil { //空链表
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id >= emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur

}

type HashTable struct {
	LinArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	//添加到那个链表
	linkNo := this.HashFun(emp.Id)
	//对应的链表
	this.LinArr[linkNo].InsertE(emp)
}

func (this *HashTable) HashFun(id int) int {
	return id % 7
}

func (this *HashTable) HList() {
	for i := 0; i < len(this.LinArr); i++ {
		this.LinArr[i].showList(i)
	}
}
func (this *EmpLink) showList(no int) {
	if this.Head == nil {
		fmt.Println("链表为空...", no)
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d id=%v 名字%v ->  ", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (this *HashTable) findById(id int) *Emp {
	linkNo := this.HashFun(id)
	fmt.Println("linkNo === ", linkNo)
	return this.LinArr[linkNo].findBId(id)
}

func (this *EmpLink) findBId(lno int) *Emp {
	cur := this.Head
	for {
		if cur == nil {
			break
		}
		if cur.Id == lno && cur != nil {
			return cur
		}
		cur = cur.Next
	}
	return cur
}
func main() {

	id := 0
	name := ""
	key := ""

	var HashTable HashTable
	for {
		fmt.Println("\t ---菜单---")
		fmt.Println("\t input 添加")
		fmt.Println("\t show 添加")
		fmt.Println("\t find 添加")
		fmt.Println("\t exit 添加")
		fmt.Println("\t请输入")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入ID")
			fmt.Scanln(&id)
			fmt.Println("输入name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			HashTable.Insert(emp)
		case "show":
			HashTable.HList()
		case "find":
			fmt.Println("输入查找的id")
			fmt.Scanln(&id)
			fmt.Println("id==", id)
			e := HashTable.findById(id)
			if e == nil {
				fmt.Println("no find")
			} else {
				fmt.Println("find ", e)
			}

		case "exit":
		default:
			fmt.Println("输入有误")
		}
	}

}
