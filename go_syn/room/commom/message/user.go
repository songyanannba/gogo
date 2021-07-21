package message

type User struct {
	//为了序列化 和反序列化成功
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json :"userStatus"` //用户状态
}
