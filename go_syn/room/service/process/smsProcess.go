package process2

import (
	"encoding/json"
	"fmt"
	"net"
	"room/commom/message"
	"room/service/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmeMes
	err := json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes 失败")
	}

	data, err := json.Marshal(mes)
	//qunfa
	for id, up := range userMgr.OnlineUsers {
		//过滤自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println(err)
	}
}
