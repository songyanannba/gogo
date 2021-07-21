package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"room/client/utils"
	"room/commom/message"
)

type UserProcess struct {
}

//写一个登陆函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//开始 定协议
	fmt.Printf("userId = %d userPwd=%s \n", userId, userPwd)

	//链接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Printf("net dail err %v\n", err)
		return
	}
	defer conn.Close()

	//准备数据
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建 LoginMes
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将我们的loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Printf("json.Marshal err %v\n", err)
		return
	}

	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err %v\n", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write err %v\n", err)
		return
	}

	fmt.Printf("客户端发送消息长度成功 %v %v\n", len(buf), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("conn.Write data err %v\n", err)
		return
	}

	//这里需要处理 服务器返回的消息

	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Printf("conn.readPkg data err %v\n", err)
		return
	}

	//反序列化
	var LoginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Date), &LoginResMes)

	if LoginResMes.Code == 200 {

		//初始化 curUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnLine

		//显示当前在线用户列表
		fmt.Println("当前的用户列表")

		for _, v := range LoginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Printf("用户id: \t%v\n", v)
			//完成 客户端 onlineUsers 初始化
			//创建user
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnLine,
			}
			onlineUsers[v] = user
		}
		//协程 如果服务器有数据推送给客户端
		go serviceProcessMes(conn)

		//显示登陆成功后的菜单
		for {
			ShowMenu()
		}

	} else {
		fmt.Printf(LoginResMes.Error)
	}

	return
}

func (this *UserProcess) Register(userId int, userPwd, UserName string) (err error) {

	//链接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Printf("net dail err %v\n", err)
		return
	}
	defer conn.Close()

	//准备数据
	var mes message.Message
	mes.Type = message.RegisterMseType

	//创建 LoginMes
	var registerMes message.RegisterMse
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = UserName

	fmt.Printf("XXX %v\v", registerMes)
	// 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Printf("json.Marshal err %v\n", err)
		return
	}

	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json.Marshal err %v\n", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	//
	err = tf.WritePkg(conn, data)
	if err != nil {
		fmt.Printf("注册WritePkg err %v\n", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Printf("注册.readPkg data err %v\n", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Date), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功 重新登录")
		os.Exit(0)
	} else {
		fmt.Printf(registerResMes.Error)
		os.Exit(0)
	}

	return
}
