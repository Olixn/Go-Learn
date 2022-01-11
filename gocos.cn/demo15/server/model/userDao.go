/**
 * @Author: Ne-21
 * @Description: 定义一个 UserDao 结构体，完成对User结构体的操作
 * @File:  userDao
 * @Version: 1.0.0
 * @Date: 2022/1/10 14:23
 */

package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gocos.cn/demo15/common/message"
)

// 在服务器启动后，初始化一个UserDao实例，将其作为全局变量

var (
	MyUserDao *UserDao
)

// 定义一个 UserDao 结构体，完成对User结构体的操作

type UserDao struct {
	pool *redis.Pool // redis 连接池
}

// 使用工厂模式，创建一个UserDao实例

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据用户id 返回一个User 实例 + err
func (u *UserDao) getUserById(conn redis.Conn, id int) (user message.User, err error) {
	// 通过指定id去redis查找用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			// users 中没有对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	// 将res反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

// 完成登录校验

func (u *UserDao) Login(userId int, userPwd string) (user message.User, err error) {
	// 从连接池取出一个链接
	conn := u.pool.Get()
	defer conn.Close()

	user, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 检验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

// 处理注册

func (u *UserDao) Register(user *message.User) (err error) {
	conn := u.pool.Get()
	defer conn.Close()
	_, err = u.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// id可注册,序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误")
		return
	}
	return
}
