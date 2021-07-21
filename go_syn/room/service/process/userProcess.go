package process2

import (
	"encoding/json"
	"fmt"
	"net"
	"room/commom/message"
	"room/service/model"
	"room/service/utils"
)

type UserProcess struct {
	//链接 ？
	Conn   net.Conn
	UserId int
}

//编写通知所有在线用户的方法

func (this *UserProcess) NotifyOthersOnlineUser(userId int) (err error) {
	//遍历在线用户 一个一个发送消息
	for id, up := range userMgr.OnlineUsers {
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)

	}
	return
}

func (this *UserProcess) NotifyMeOnline(userId int) (err error) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOffline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Printf("NotifyMeOnline json.Marshal err== %v\n", err)
		return
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("NotifyMeOnline json.Marshal mes err== %v\n", err)
		return
	}

	//创建Transfer 实例 发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline WritePkg err = ", err)
		return
	}
	return
}

func (this *UserProcess) ServiceProcessRegister(msg *message.Message) (err error) {
	var registerMes message.RegisterMse

	err = json.Unmarshal([]byte(msg.Date), &registerMes)
	if err != nil {
		fmt.Printf("ServiceProcessRegister 反序列化错误 %v\n", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterMseType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXIST.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("登陆成功")
	}

	dat, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Printf("ServiceProcessRegister 序列化错误 %v\n", err)
		return
	}

	resMes.Date = string(dat)

	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Printf("ServiceProcessRegister 序列化错误 %v\n", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

//函数 处理登陆的请求
func (this *UserProcess) ServiceProcesslogin(msg *message.Message) (err error) {
	//从mes 里面 获取data message 反序列化LogMes
	var loginMes message.LoginMes

	err = json.Unmarshal([]byte(msg.Date), &loginMes)
	if err != nil {
		fmt.Printf("serviceProcesslogin 反序列化错误 %v\n", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginMesType

	//在声明一个 LoginResMes
	var LoginResMes message.LoginResMes

	//如果用户id =100 密码 123456 合法

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		//不合法
		if err == model.ERROR_USER_NOTEXISTS {
			LoginResMes.Code = 500
			LoginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			LoginResMes.Code = 403
			LoginResMes.Error = err.Error()
		} else {
			LoginResMes.Code = 505
			LoginResMes.Error = "不确定的错误\n"
		}

	} else {
		LoginResMes.Code = 200
		//把登陆成功的用户 放入到usermgr里面

		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其他在线用户 我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前的user_id 放入loginResMes
		for id, _ := range userMgr.OnlineUsers {
			LoginResMes.UserIds = append(LoginResMes.UserIds, id)
		}

		fmt.Printf("xxx登陆成功 %v \n", user)
	}

	dat, err := json.Marshal(LoginResMes)

	if err != nil {
		fmt.Printf("LoginResMes 序列化错误 %v\n", err)
		return
	}
	// 将 dat 负值给resMes
	resMes.Date = string(dat)

	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Printf("resMes 序列化错误 %v\n", err)
		return
	}

	//将发送的消息封装

	//分层需要 先创建 transfer

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)

	return
}
