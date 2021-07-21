package service

import (
	"fmt"
	"kehuguanli/model"
)

//完成怼customer的操作
//增删改查
type CustomerService struct {
	customers []model.Customer

	//定义一个字段，有多少客户 ;可以作为新客户的ID
	customerNum int
}

func NewCustomerService() *CustomerService {
	//初始化数据
	customerService := &CustomerService{}
	customerService.customerNum = 1

	customer := model.NewCustomer(1, "张三", "男", 20, "11121212", "zhangsan@qq.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers

}

//添加客户到切片中
func (this *CustomerService) Add(customer model.Customer) bool {

	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Update(customer model.Customer) bool {

	index := this.findID(customer.Id)
	if index == -1 {
		fmt.Println("当前的index")
		fmt.Println(index)
		return false
	}
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == index {
			this.customers[i].Name = customer.Name
			this.customers[i].Age = customer.Age
			this.customers[i].Gender = customer.Gender
			this.customers[i].Phone = customer.Phone
			this.customers[i].Email = customer.Email
		}
	}

	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.findID(id)
	if index == -1 {
		fmt.Println("当前的indexindex")
		fmt.Println(index)
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) findID(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = id
		}
	}
	return index
}
