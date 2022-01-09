/**
 * @Author: Ne-21
 * @Description: 消息结构体
 * @File:  message
 * @Version: 1.0.0
 * @Date: 2022/1/9 16:43
 */

package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
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
	Code  int    `json:"code"`  // 返回状态码 500 未注册 200 登录成功
	Error string `json:"error"` // 返回的错误信息
}

type RegisterMes struct {
}
