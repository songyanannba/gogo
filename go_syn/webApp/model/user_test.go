package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户...")
	//user := &User{}
	//user.AddUser()
	//user.AddUser2()
}

func TestGetUserById(t *testing.T) {
	fmt.Println("查询一条用户数据")
	user := User{
		Id: 4,
	}
	u, _ := user.GetUserById()
	fmt.Println("user信息", u)
}

func TestGetUsers(t *testing.T) {
	fmt.Println("查询全部用户数据")
	user := &User{}
	us, _ := user.GetUsers()

	for k, v := range us {
		fmt.Println(k)
		fmt.Println(v)
	}
}
