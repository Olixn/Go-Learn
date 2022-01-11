/**
 * @Author: Ne-21
 * @Description:
 * @File:  smsProcess
 * @Version: 1.0.0
 * @Date: 2022/1/11 17:12
 */

package controller

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
)

type SmsProcess struct {
}

// 发送群聊消息

func (s *SmsProcess) SendGroupMes(content string) (err error) {
	// 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return err
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return err
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err = ", err)
		return err
	}
	return
}
