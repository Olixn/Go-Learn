/**
 * @Author: Ne-21
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/9 10:46
 */

package main

import (
	"fmt"
	"gocos.cn/demo15/client/controller"
)

var userId int
var userPwd string
var userName string

func main() {
	// 接受用户选择
	var key int
	// 判断是否继续显示菜单
	var loop = true
	for loop {
		fmt.Println("---------欢迎登录多人聊天系统---------")
		fmt.Println("\t\t\t1 登录聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			// 用户要登录
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			// 登录
			up := &controller.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字：")
			fmt.Scanf("%s\n", &userName)
			// 注册
			up := &controller.UserProcess{}
			up.Register(userId, userPwd, userName)
			// loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，重新输入")
		}
	}
}
