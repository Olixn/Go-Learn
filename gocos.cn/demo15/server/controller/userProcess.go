/**
 * @Author: Ne-21
 * @Description: 处理用户相关
 * @File:  userProcess
 * @Version: 1.0.0
 * @Date: 2022/1/9 19:59
 */

package controller

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
	"gocos.cn/demo15/server/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	// 标识用户
	UserId int
}

// 通知所有在线用户的方法

func (u *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历在线用户切片，逐一发送
	for id, up := range MyUserMgr.OnlineUsers {
		if id == userId {
			continue
		}
		// 开始推送
		up.NotifyMeOnline(userId)
	}
}

func (u *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: u.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err = ", err)
		return
	}
}

// 编写一个ServerProcessLogin 函数 来处理登录请求

func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 登录处理核心代码
	// 从 mes 取出 mes.Data 并进行反序列化 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 在声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 操作redis数据库
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 // 用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
		// 把该登陆成功的用户放入在线认识管理UserMgr中
		// 将登陆成功userID赋给u
		u.UserId = loginMes.UserId
		MyUserMgr.AddOnlineUser(u)
		// 主动推送自己的登录信息个其他人
		u.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户id放入loginResMes.UsersId
		for id, _ := range MyUserMgr.OnlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

	}

	// 序列化 loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 将 data 赋值给resMes.Data
	resMes.Data = string(data)

	// 序列化 resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 发送 data 将其封装到writePkg函数中
	// err = utils.WritePkg(conn, data)
	// 分层重写
	tr := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tr.WritePkg(data)
	return
}

func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 从 mes 取出 mes.Data 并进行反序列化 RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	// 在声明一个 RegisterResMes
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 400
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "服务器内部错误"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User, "注册成功")
	}

	// 序列化 registerResMes
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 将 data 赋值给resMes.Data
	resMes.Data = string(data)

	// 序列化 resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 发送 data
	tr := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tr.WritePkg(data)
	return
}
