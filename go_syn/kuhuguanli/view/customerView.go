package main

import (
	"fmt"
	"kehuguanli/model"
	"kehuguanli/service"
)

type customerView struct {
	key  string //接受用户输入
	loop bool   //是否循环展示菜单

	customerService *service.CustomerService
}

//展示客户信息
func (this *customerView) list() {
	customers := this.customerService.List()
	//展示
	fmt.Println("\n\n\n----客户列表----")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----客户列表完成----")
}

func (this *customerView) add() {
	//得到用户输入信息 构建数据 调用添加
	fmt.Println("【添加客户】")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	//构建customer
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}
}

func (this *customerView) update() {
	//得到用户输入信息 构建数据
	fmt.Println("【修改客户】")

	fmt.Println("编号")
	id := 0
	fmt.Scanln(&id)

	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	//构建customer
	customer := model.NewCustomer(id, name, gender, age, phone, email)

	if this.customerService.Update(customer) {
		fmt.Println("修改完成")
	} else {
		fmt.Println("修改失败")
	}

}

func (this *customerView) delete() {
	fmt.Println("【输入删除用户编号】")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("放弃删除")
		return
	}
	fmt.Println("确认输入Y或者y")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//查找ID对应的用户 -1没有相关客户
		if !this.customerService.Delete(id) {
			fmt.Println("输入的ID有误")
		} else {
			fmt.Println("删除成功")
			this.list()
		}
	} else {
		fmt.Println("确认错误")
		return
	}

}

func (this *customerView) mainView() {
	for {
		fmt.Println("\n\n\n---客户管理---")
		fmt.Println("  1添加客户")
		fmt.Println("  2修改客户")
		fmt.Println("  3删除客户")
		fmt.Println("  4客户列表")
		fmt.Println("  5推 出  ")
		fmt.Println("  请选择1--5")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添加客户")
			this.add()
		case "2":
			fmt.Println("修改客户")
			this.update()
		case "3":
			fmt.Println("删除客户")
			this.delete()
		case "4":
			this.list()
		case "5":
			//this.loop = false
			this.exit()
		default:
			fmt.Println("输入有误，请选择1--5")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出程序")
}

func (this *customerView) exit() {
	fmt.Println("确定要推出吗 y/n")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("输入有误")
		fmt.Println(this.key)
	}

	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()

	customerView.mainView()
}
