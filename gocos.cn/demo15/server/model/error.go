/**
 * @Author: Ne-21
 * @Description: 定义错误信息
 * @File:  error
 * @Version: 1.0.0
 * @Date: 2022/1/10 14:23
 */

package model

import "errors"

// 根据业务逻辑需要，定义一些错误

var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在。")
	ERROR_USER_EXISTS    = errors.New("用户已存在。")
	ERROR_USER_PWD       = errors.New("密码不正确")
)
