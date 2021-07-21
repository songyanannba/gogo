package model

import (
	"fmt"
	"webApp/utils"
)

type User struct {
	Id        int
	User_name string
	Pwd       string
}

func (user *User) AddUser() error {
	//sql语句
	fmt.Println("111")
	sqlStr := "INSERT INTO users(user_name,pwd) VALUES(? ,?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf(" Prepare 异常%v\n ", err)
		return err
	}
	_, err = stmt.Exec("admin", "123456")
	if err != nil {
		fmt.Printf("Exec 异常 %v \n", err)
		return err
	}
	return nil
}

func (user *User) AddUser2() error {
	//sql语句
	fmt.Println("222")
	sqlStr := "INSERT INTO users(user_name,pwd) VALUES(? ,?)"
	_, err := utils.Db.Exec(sqlStr, "admin2", "666")
	if err != nil {
		fmt.Println("异常", err)
		return err
	}
	return nil
}

func (user *User) GetUserById() (*User, error) {
	sqlstr := "select * from users where id = ?"
	row := utils.Db.QueryRow(sqlstr, user.Id)
	//声明
	var id int
	var user_name string
	var pwd string

	err := row.Scan(&id, &user_name, &pwd)

	if err != nil {
		return nil, err
	}
	u := &User{
		Id:        id,
		User_name: user_name,
		Pwd:       pwd,
	}
	return u, nil
}

func (user *User) GetUsers() ([]*User, error) {
	sqlstr := "select * from users "
	row, _ := utils.Db.Query(sqlstr)

	var users []*User
	for row.Next() {
		var id int
		var user_name string
		var pwd string
		err := row.Scan(&id, &user_name, &pwd)
		if err != nil {
			return nil, err
		}
		u := &User{
			Id:        id,
			User_name: user_name,
			Pwd:       pwd,
		}
		users = append(users, u)
	}
	return users, nil

}
