package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMseType         = "RegisterMse"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnLine = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Date string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"`  //返回的状态吗 500 用户未注册 200 登陆成功
	Error   string `json:"error"` //
	UserIds []int
}

type RegisterMse struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` //400 已经存在 200 注册成功
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmeMes struct {
	Content string `json :"content"`
	User
}
