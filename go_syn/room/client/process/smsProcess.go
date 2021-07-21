package process

import (
	"encoding/json"
	"fmt"
	"room/client/utils"
	"room/commom/message"
)

type SmsProcess struct {
}

//发送群聊
func (this *SmsProcess) SenGroupMes(content string) (err error) {
	//创建mes
	var mes message.Message
	mes.Type = message.SmsMesType

	var SmeMes message.SmeMes
	SmeMes.Content = content
	SmeMes.UserId = CurUser.UserId
	SmeMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(SmeMes)
	if err != nil {
		fmt.Println(err)
	}
	mes.Date = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(CurUser.Conn, data)

	return
}
