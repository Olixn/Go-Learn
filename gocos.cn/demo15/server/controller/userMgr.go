/**
 * @Author: Ne-21
 * @Description:
 * @File:  userMgr
 * @Version: 1.0.0
 * @Date: 2022/1/11 13:10
 */

package controller

import (
	"fmt"
)

// 声明一个全局的在线人数管理,并在启动时初始化

var MyUserMgr *UserMgr

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

// 添加

func (u *UserMgr) AddOnlineUser(up *UserProcess) {
	u.OnlineUsers[up.UserId] = up
}

// 删除

func (u *UserMgr) DeleteOnlineUser(userId int) {
	delete(u.OnlineUsers, userId)
}

// 查询

func (u *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return u.OnlineUsers
}

func (u *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := u.OnlineUsers[userId]
	if !ok {
		// 用户当前不在线
		err = fmt.Errorf("用户 %d 不在线", userId)
		return
	}
	return
}
