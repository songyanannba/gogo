package model

import (
	"encoding/json"
	"fmt"
	"room/commom/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

//定义userDao 完成多user的各种操作
type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//根据用户ID 返回user 实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定的id 去redis 查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	fmt.Printf("REDIS %v\n", res)
	if err != nil {
		if err == redis.ErrNil {
			//在user的哈希中 没有这个人
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//把res 反序列化成user 实例
	user = &User{}
	err = json.Unmarshal([]byte(res), user)

	if err != nil {
		fmt.Printf("getUserById json.Unmarshal err == %v \n", err)
		return
	}
	return
}

//完成对用户的验证
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao的链接中 获取一根链接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {

		return
	}
	//证明拿到user
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	//先从UserDao的链接中 获取一根链接
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return
	}

	//还没有注册redis 完成注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Register 序列化错误")
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Printf("Register 注册用户错误err %v\n", err)
	}
	return
}
