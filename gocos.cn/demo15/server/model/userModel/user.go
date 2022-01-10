/**
 * @Author: Ne-21
 * @Description: 用户的结构体
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/1/10 14:23
 */

package userModel

// 定义一个用户的结构体

type User struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
