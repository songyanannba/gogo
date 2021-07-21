package process

import (
	"encoding/json"
	"fmt"
	"room/commom/message"
)

func outputGroupMes(mes *message.Message) {
	//
	var smsMes message.SmeMes
	err := json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	info := fmt.Sprintf("用户id :\t %d 对大家说: \t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)

}
