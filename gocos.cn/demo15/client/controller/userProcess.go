/**
 * @Author: Ne-21
 * @Description: 客户端处理用户相关功能
 * @File:  userProcess
 * @Version: 1.0.0
 * @Date: 2022/1/10 13:10
 */

package controller

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
	"net"
)

type UserProcess struct {
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	// 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 准备通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	// 创建一个 LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	// 把 data 赋给 mes.Data
	mes.Data = string(data)

	// 将 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}
	// 发送数据
	// 创建一个Transfer 实例
	tr := &utils.Transfer{
		Conn: conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg err = ", err)
		return
	}

	// 处理服务器端返回的消息
	mes, err = tr.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg err = ", err)
		return
	}
	// mes 的 Data 反序列化 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登陆成功")

		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// 显示当前在线用户列表
		fmt.Println("当前在线列表：")
		for _, v := range loginResMes.UsersId {
			if v == loginMes.UserId {
				continue
			}
			fmt.Println("用户id = ", v)

			// 完成对 客户端的 OnlineUsers 的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			OnlineUsers[v] = user
		}
		fmt.Println("---------------------------")
		// 在客户端启用协程，保持服务器通讯 服务端推送消息-接收-显示在客户端
		go ProcessServerMes(conn)
		// 调用登陆成功的菜单[循环]
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (u *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 准备通过conn发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType
	// 创建一个 LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 把 data 赋给 mes.Data
	mes.Data = string(data)

	// 将 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}
	// 发送数据
	// 创建一个Transfer 实例
	tr := &utils.Transfer{
		Conn: conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("utils.WritePkg err = ", err)
		return
	}

	// 处理服务器端返回的消息
	mes, err = tr.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg err = ", err)
		return
	}
	// mes 的 Data 反序列化 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登陆")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}
