/**
 * @Author: Ne-21
 * @Description:
 * @File:  smsMgr
 * @Version: 1.0.0
 * @Date: 2022/1/11 19:11
 */

package controller

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
)

func OutPutGroupMes(mes *message.Message) {
	var smsResMes message.SmsResMes
	err := json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 显示信息
	info, _ := fmt.Printf("用户id：\t%d 对大家说：\t%s", smsResMes.UserId, smsResMes.Content)
	fmt.Println(info)
	fmt.Println()
}
