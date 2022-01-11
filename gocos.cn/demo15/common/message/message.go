/**
 * @Author: Ne-21
 * @Description: 消息结构体
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2022/1/9 16:43
 */

package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SmsResMesType           = "SmsResMes"
)

// 定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 定义两种消息

type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code    int    `json:"code"`    // 返回状态码 500 未注册 200 登录成功
	Error   string `json:"error"`   // 返回的错误信息
	UsersId []int  `json:"usersId"` // 增加字段，保持在线用户id
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 已占用 200 注册成功
	Error string `json:"error"` // 返回的错误信息
}

// 配合服务器端的推送

type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

// 增加一个SmsMes

type SmsMes struct {
	User           // 匿名结构体
	Content string `json:"content,omitempty"` // 内容
}

type SmsResMes struct {
	User           // 匿名结构体
	Content string `json:"content,omitempty"` // 内容
}
