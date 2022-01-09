/**
 * @Author: Ne-21
 * @Description: 处理用户相关
 * @File:  userProcess
 * @Version: 1.0.0
 * @Date: 2022/1/9 19:59
 */

package userProcess

import (
	"encoding/json"
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
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

	// 如果用户id 100 密码 123456 否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500 // 用户不存在
		loginResMes.Error = "该用户不存在，请先注册"
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
