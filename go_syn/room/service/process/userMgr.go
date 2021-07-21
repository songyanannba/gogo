package process2

import "fmt"

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

var (
	userMgr *UserMgr
)

//完成初始化
func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//添加在线用户

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.OnlineUsers[up.UserId] = up
}

func (this *UserMgr) DelonLineUser(UserId int) {
	delete(this.OnlineUsers, UserId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUsers
}

func (this *UserMgr) GetOnlineUserById(UserId int) (up *UserProcess, err error) {
	//从map里面取值
	up, ok := this.OnlineUsers[UserId]
	if !ok { //当前用户不在线
		err = fmt.Errorf("用户id 不存在")
		return
	}
	return
}
