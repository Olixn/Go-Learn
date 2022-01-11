/**
 * @Author: Ne-21
 * @Description:
 * @File:  smsProcess
 * @Version: 1.0.0
 * @Date: 2022/1/11 18:49
 */

package controller

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
	"net"
)

type SmsProcess struct {
}

// 转发消息

func (s *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 更改type SmsMes -》 SmsResMes
	mes.Type = message.SmsResMesType
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 遍历OnlineUsers，转发消息
	for id, up := range MyUserMgr.OnlineUsers {
		if id == smsMes.UserId {
			continue
		}
		s.SendEachOnlineUser(data, up.Conn)
	}
}

func (s *SmsProcess) SendEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendEachOnlineUser err = ", err)
		return
	}
}
