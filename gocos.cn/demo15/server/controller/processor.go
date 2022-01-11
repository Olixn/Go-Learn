/**
 * @Author: Ne-21
 * @Description: 总控制器
 * @File:  processor
 * @Version: 1.0.0
 * @Date: 2022/1/9 19:58
 */

package controller

import (
	"fmt"
	"gocos.cn/demo15/common/message"
	"gocos.cn/demo15/common/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes 函数 来处理不同消息

func (p *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		up := &UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 处理短消息,创建一个smsProcess实例来转发消息
		sp := &SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理......")
	}
	return
}

func (p *Processor) Process() (err error) {
	// 循环读客户端发送的信息
	for {
		// 将读取数据包封装成函数readPkg(),返回Message，Err
		// 创建一个Transfer 实例 完成读包任务
		tr := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tr.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了，关闭链接")
				return err
			} else {
				fmt.Println("readPkg err = ", err)
				return err
			}
		}
		fmt.Println("mes = ", mes)
		err = p.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err = ", err)
			return err
		}

	}
}
