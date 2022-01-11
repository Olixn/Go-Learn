/**
 * @Author: Ne-21
 * @Description:
 * @File:  modelDao
 * @Version: 1.0.0
 * @Date: 2022/1/10 15:18
 */

package initService

import (
	"gocos.cn/demo15/server/model"
)

func InitUserDao() {
	// pool 是全局变量，并且通过InitRedisPool初始化，所以本函数调用要在其 之后
	model.MyUserDao = model.NewUserDao(pool)
}
