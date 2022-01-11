/**
 * @Author: Ne-21
 * @Description:
 * @File:  userMgr
 * @Version: 1.0.0
 * @Date: 2022/1/11 16:12
 */

package controller

import (
	"fmt"
	"gocos.cn/demo15/client/model"
	"gocos.cn/demo15/common/message"
)

// 客户端维护的map

var OnlineUsers = make(map[int]*message.User, 10)
var CurUser model.CurUser // 在用户登录成功后，对其初始化

// 处理返回的NotifyUserStatusMes

func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := OnlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	OnlineUsers[notifyUserStatusMes.UserId] = user
	outPutOnlineUser()
}

// 显示当前在线用户

func outPutOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range OnlineUsers {
		fmt.Println("用户id = ", id)
	}
}
