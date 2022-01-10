/**
 * @Author: Ne-21
 * @Description: 显示二级服务菜单、保持通讯
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2022/1/10 13:13
 */

package server

import (
	"fmt"
	"gocos.cn/demo15/common/utils"
	"net"
	"os"
)

// 显示登陆成功后的界面

func ShowMenu() {
	fmt.Println("-------恭喜xx登陆成功-------")
	fmt.Println("1 显示在线用户列表")
	fmt.Println("2 发送消息")
	fmt.Println("3 消息列表")
	fmt.Println("4 退出系统")
	fmt.Println("请选择（1-4）：")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入。")
	}
}

// 和服务器端保持通信

func ProcessServerMes(conn net.Conn) {
	// 创建一个Transfer实例，不停地读服务器发送的消息
	//
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}
		// 如果读取到消息，进行下一步处理
		fmt.Println("mes = ", mes)
	}
}
